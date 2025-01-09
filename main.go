package main

import (
	"net/http"

	"github.com/andrew-lawlor/cc-simple-auth-server/auth"
	"github.com/andrew-lawlor/cc-simple-auth-server/token"
)

func main() {
	// Init tokens.
	token.LoadTokens()
	// Set up routes.
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", auth.Register)
	mux.HandleFunc("POST /login", auth.Login)
	// Start server.
	http.ListenAndServe("localhost:3010", mux)
}
