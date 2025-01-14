package main

import (
	"fmt"
	"net/http"
)

func useServerMode(config Config) {
	p := 3000

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("config: " + config.mode))
	})

	fmt.Printf("Server listening on http://localhost:%d\n", p)
	http.ListenAndServe(fmt.Sprintf(":%d", p), mux)
}
