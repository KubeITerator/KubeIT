package gateway

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"io"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/user"
	"net/http"
	"time"
)

type TempKey struct {
	RToken    string
	ExpiresAt time.Time
}

type Gateway struct {
	gwmux         *runtime.ServeMux
	tempkeyaccess map[string]TempKey
	db            *db.Database
}

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (gw *Gateway) Init(clientid, secret string, database *db.Database) {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests

	gw.db = database
	gw.tempkeyaccess = make(map[string]TempKey)
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "init",
			"topic": "grpc_gateway",
			"key":   "dial_grpc",
		}).Fatal("grpc dial failed: " + err.Error())
	}

	gw.gwmux = runtime.NewServeMux()

	err = gw.HandleAuth(context.Background(), clientid, secret)

	if err != nil {
		log.WithFields(log.Fields{
			"stage": "init",
			"topic": "handle_auth",
			"key":   "route_setup_failed",
		}).Fatal("handle auth setup_routes failed: " + err.Error())
	}

	err = user.RegisterUserManagerHandler(context.Background(), gw.gwmux, conn)
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "init",
			"topic": "register_gateway",
			"key":   "user_manager_gateway",
		}).Fatal("Failed to register user manager handler: " + err.Error())
	}

	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gw.gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8091")
	log.Fatalln(gwServer.ListenAndServe())
}

func (gw *Gateway) randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (gw *Gateway) setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func (gw *Gateway) CreateTempKey(refreshToken string, claims *Claims) (string, error) {

	rnd, err := gw.randString(10)
	if err != nil {
		return "", err
	}

	if !gw.db.UserExists(claims.Sub) {
		u := user.User{
			Id:           "",
			Sub:          claims.Sub,
			Name:         claims.Name,
			Email:        claims.Email,
			Admin:        false,
			Tokens:       []*user.Token{},
			GPermissions: []*user.GroupPermission{},
		}
		_, err := gw.db.AddUser(&u)
		if err != nil {
			return "", err
		}
	}

	// Tempkeys expire after 30 seconds
	key := TempKey{
		RToken:    refreshToken,
		ExpiresAt: time.Now().Add(30 * time.Second),
	}

	gw.tempkeyaccess[rnd] = key
	return rnd, nil
}

func (gw *Gateway) HandleAuth(ctx context.Context, clientid, secret string) error {

	provider, err := oidc.NewProvider(ctx, "http://localhost:8090/auth/realms/kubeit-test")
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "handle_auth",
			"key":   "new_provider",
		}).Warn("Provider init failed")
		return err
	}
	oidcConfig := &oidc.Config{
		ClientID: clientid,
	}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientid,
		ClientSecret: secret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:8091/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	err = gw.gwmux.HandlePath("GET", "/auth/login", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		state, err := gw.randString(16)
		if err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "login_route",
				"key":   "state",
			}).Warn("Random state creation failed: " + err.Error())
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		nonce, err := gw.randString(16)
		if err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "login_route",
				"key":   "nonce",
			}).Warn("Random nonce creation failed: " + err.Error())
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		gw.setCallbackCookie(w, r, "state", state)
		gw.setCallbackCookie(w, r, "nonce", nonce)

		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
	})

	if err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "login_route",
			"key":   "gmux_error",
		}).Warn("Gmux error in handling route: " + err.Error())
		return err
	}

	err = gw.gwmux.HandlePath("OPTIONS", "/auth/retrieve", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	})

	if err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "auth_retrieve",
			"key":   "gmux_error_options",
		}).Warn("Gmux error in handling route: " + err.Error())
		return err
	}

	err = gw.gwmux.HandlePath("GET", "/auth/retrieve", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

		query := r.URL.Query()

		qid := query.Get("id")
		temptoken := gw.tempkeyaccess[qid]

		if time.Now().Before(temptoken.ExpiresAt) {
			rtoken := struct {
				Token string
			}{temptoken.RToken}

			data, err := json.MarshalIndent(rtoken, "", "    ")
			if err != nil {

				log.WithFields(log.Fields{
					"stage": "auth",
					"topic": "auth_retrieve",
					"key":   "marshal_token",
				}).Warn("Token marshalling failed: " + err.Error())
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}

			// TODO: Specify a distinct dns name as origin
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			delete(gw.tempkeyaccess, qid)
			return
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusGone)
			delete(gw.tempkeyaccess, qid)
			return
		}
	})

	if err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "auth_retrieve",
			"key":   "gmux_error_post",
		}).Warn("Gmux error in handling route: " + err.Error())
		return err
	}

	err = gw.gwmux.HandlePath("GET", "/auth/callback", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

		state, err := r.Cookie("state")

		if err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "state_not_found",
			}).Warn("state not found")
			http.Error(w, "state not found", http.StatusBadRequest)
			return
		}
		if r.URL.Query().Get("state") != state.Value {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "state_mismatch",
			}).Warn("state mismatched")
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "token_exchange",
			}).Warn("token exchange failed: " + err.Error())
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "id_token_not_found",
			}).Warn("id token could not be found")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "verify_id_token",
			}).Warn("failed to verify id_token")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		nonce, err := r.Cookie("nonce")
		if err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "nonce_not_found",
			}).Warn("nonce could not be found")
			http.Error(w, "nonce not found", http.StatusBadRequest)
			return
		}
		if idToken.Nonce != nonce.Value {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "nonce_mismatched",
			}).Warn("nonce could not be matched")
			http.Error(w, "nonce did not match", http.StatusBadRequest)
			return
		}

		claims := &Claims{}

		if err := idToken.Claims(claims); err != nil {
			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "id_token_match",
			}).Warn("Id token could not be mapped to claims")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		key, err := gw.CreateTempKey(oauth2Token.RefreshToken, claims)

		if err != nil {

			log.WithFields(log.Fields{
				"stage": "auth",
				"topic": "auth_callback",
				"key":   "temp_key_creation_failed",
			}).Warn("temp key creation failed: " + err.Error())
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("http://127.0.0.1:3000/callback?id=%s", key), http.StatusFound)
	})

	return nil

}
