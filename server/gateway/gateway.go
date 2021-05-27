package gateway

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"io"
	"kubeIT/pkg/grpc/user"
	"log"
	"net/http"
	"os"
	"time"
)

type TempKey struct {
	RToken    string
	ExpiresAt time.Time
}

type Gateway struct {
	gwmux         *runtime.ServeMux
	tempkeyaccess map[string]TempKey
}

type TestClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (gw *Gateway) Init(clientid, secret string) {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests

	gw.tempkeyaccess = make(map[string]TempKey)
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gw.gwmux = runtime.NewServeMux()

	gw.HandleAuth(context.Background(), clientid, secret)

	err = user.RegisterUserManagerHandler(context.Background(), gw.gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
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

func (gw *Gateway) CreateTempKey(refreshToken string) (string, error) {

	rnd, err := gw.randString(10)
	if err != nil {
		return "", err
	}

	// Tempkeys expire after 30 seconds
	key := TempKey{
		RToken:    refreshToken,
		ExpiresAt: time.Now().Add(30 * time.Second),
	}

	gw.tempkeyaccess[rnd] = key
	return rnd, nil
}

func (gw *Gateway) HandleAuth(ctx context.Context, clientid, secret string) {

	provider, err := oidc.NewProvider(ctx, "http://localhost:8090/auth/realms/kubeit-test")
	if err != nil {
		log.Fatal(err)
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

	err = gw.gwmux.HandlePath("GET", "/login", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		state, err := gw.randString(16)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		nonce, err := gw.randString(16)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		gw.setCallbackCookie(w, r, "state", state)
		gw.setCallbackCookie(w, r, "nonce", nonce)

		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
	})

	if err != nil {
		fmt.Println("Error in gwmux: " + err.Error())
		os.Exit(2)
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
				fmt.Println("Error in gwmux: " + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// TODO: Specify a distinct dns name as origin
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		} else {
			w.WriteHeader(http.StatusGone)
		}
		delete(gw.tempkeyaccess, qid)
	})

	if err != nil {
		fmt.Println("Error in gwmux: " + err.Error())
		os.Exit(2)
	}

	err = gw.gwmux.HandlePath("GET", "/auth/callback", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		state, err := r.Cookie("state")

		fmt.Println(r.Cookies())
		if err != nil {
			http.Error(w, "state not found", http.StatusBadRequest)
			return
		}
		if r.URL.Query().Get("state") != state.Value {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		nonce, err := r.Cookie("nonce")
		if err != nil {
			http.Error(w, "nonce not found", http.StatusBadRequest)
			return
		}
		if idToken.Nonce != nonce.Value {
			http.Error(w, "nonce did not match", http.StatusBadRequest)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *TestClaims // ID Token payload is just JSON.
		}{oauth2Token, &TestClaims{}}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//data, err := json.MarshalIndent(resp, "", "    ")
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		fmt.Println("returned Token")

		key, err := gw.CreateTempKey(resp.OAuth2Token.RefreshToken)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("http://127.0.0.1:3000/callback?id=%s", key), http.StatusFound)
	})

}
