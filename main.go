package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir("."))
	mux.Handle("/", fileHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
