package helpers

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type Authorizer struct {
	Provider   *oidc.Provider
	OidcConfig *oidc.Config
	Verifier   *oidc.IDTokenVerifier
	Config     oauth2.Config
}

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (a *Authorizer) Init(clientid, secret string) {
	var err error
	a.Provider, err = oidc.NewProvider(context.Background(), "http://localhost:8090/auth/realms/kubeit-test")
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "oicd_provider",
			"key":   "new_provider",
		}).Fatal("Provider init failed")
	}
	a.OidcConfig = &oidc.Config{
		ClientID: clientid,
	}
	a.Verifier = a.Provider.Verifier(a.OidcConfig)

	a.Config = oauth2.Config{
		ClientID:     clientid,
		ClientSecret: secret,
		Endpoint:     a.Provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:8091/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

}

func (a *Authorizer) Verify(ctx context.Context, token *oauth2.Token) (bool, string, *Claims) {

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "auth_callback",
			"key":   "id_token_not_found",
		}).Warn("id token could not be found")
		return false, "", nil
	}
	idToken, err := a.Verifier.Verify(ctx, rawIDToken)
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "auth_callback",
			"key":   "verify_id_token",
		}).Warn("failed to verify id_token")
		return false, "", nil
	}

	claims := &Claims{}

	if err := idToken.Claims(claims); err != nil {
		log.WithFields(log.Fields{
			"stage": "auth",
			"topic": "auth_callback",
			"key":   "id_token_match",
		}).Warn("Id token could not be mapped to claims")
		return false, "", nil
	}

	return true, idToken.Nonce, claims

}
