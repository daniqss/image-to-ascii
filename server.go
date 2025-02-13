package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
)

func useServerMode(config Config) {
	p := 3000

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>oumaiga</h1>"))
	})
	mux.HandleFunc("POST /api/v1/", func(w http.ResponseWriter, r *http.Request) {
		handleImageUploaded(w, r, config)
	})

	fmt.Printf("Server listening on http://localhost:%d\n", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", p), mux))
}

func handleImageUploaded(w http.ResponseWriter, r *http.Request, config Config) {
	// parse form with 10MB size limit
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error": "Failed to parse form: %v"}`, err), http.StatusBadRequest)
		return
	}

	// get file from form
	file, _, err := r.FormFile("image")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error": "Failed to get file: %v"}`, err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// decode file as an image
	img, _, err := image.Decode(file)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error": "Failed to decode image: %v"}`, err), http.StatusBadRequest)
		return
	}

	// create ascii generator
	ascii := Ascii{
		config: config,
		img:    img,
	}
	w.Header().Set("Content-Type", "image/png")

	var writer io.Writer = w
	if err := ascii.generateAscii(&writer); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error": "Failed to generate ascii: %v"}`, err), http.StatusInternalServerError)
		return
	}
}
