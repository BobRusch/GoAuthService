package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	// OauthFacebook
	mux.HandleFunc("/auth/facebook/login", oauthFacebookLogin)
	mux.HandleFunc("/auth/facebook/callback", oauthFacebookCallback)

	return mux
}
