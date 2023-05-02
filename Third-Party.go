package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
		
)

func main() {
	// Create the authentication middleware handler
	authHandler := httpauth.SimpleBasicAuth("RayMC19", "rayray19")

	// Create the final handler
	finalHandler := http.HandlerFunc(final)

	// Chain the middleware and final handlers together
	chain := authHandler(finalHandler)

	// Create the server mux and register the handler chain
	mux := http.NewServeMux()
	mux.Handle("/", chain)

	// Start the server
	log.Print("Listening on :2023...")
	err := http.ListenAndServe(":2023", mux)
	log.Fatal(err)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Just Chill"))
}
