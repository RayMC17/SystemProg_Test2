package main

import (
	"log"
	"mime"
	"net/http"
)

// JSONHandler is a middleware that checks if the request has the "application/json" Content-Type header
func JSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		//tried to use a switch statement here but it gave me a different output not sure why.
		if contentType != "" {
			// ParseMediaType returns the media type and any parameters
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				// Return a 400 Bad Request if the Content-Type header is malformed
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			// Return a 415 Unsupported Media Type if the Content-Type is not "application/json"
			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// final is a simple handler that returns "OK"
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's fine, it's working"))
}

func main() {
	mux := http.NewServeMux()

	// Wrap the final handler with the enforceJSONHandler middleware
	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", JSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
