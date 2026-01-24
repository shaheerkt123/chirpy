package main

import (
	"net/http"
)

func main() {
	// 1. Create the multiplexer
	mux := http.NewServeMux()

	// 2. Set up the file server handler
	// We point it to the current directory (".")
	fileHandler := http.FileServer(http.Dir("."))

	// 3. Register the file server at /app/
	// We use http.StripPrefix so the file server doesn't look for an actual "app" folder
	mux.Handle("/app/", http.StripPrefix("/app", fileHandler))
	// 4. Register the readiness endpoint using mux.HandleFunc
	// Note: using "/healthz" (no trailing slash) for an exact match
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 5. Configure and start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
