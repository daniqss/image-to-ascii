package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func useServerMode(config Config) {
	p := 3000

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/", func(w http.ResponseWriter, r *http.Request) {
		handleImageUploaded(w, r, config)
	})
	mux.HandleFunc("GET /api/v1/fonts", handleFonts)

	fmt.Printf("Server listening on http://localhost:%d\n", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", p), mux))
}

func handleFonts(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Read files from the fonts directory
	files, err := os.ReadDir("./fonts")
	if err != nil {
		log.Printf("Error reading fonts directory: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to read fonts directory"}`))
		return
	}

	// Extract font names without the .ttf extension
	var fontNames []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".ttf") {
			// Remove .ttf extension
			fontName := strings.TrimSuffix(file.Name(), ".ttf")
			fontNames = append(fontNames, fontName)
		}
	}

	// Create JSON response
	response, err := json.Marshal(map[string][]string{"fonts": fontNames})
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to generate response"}`))
		return
	}

	w.Write(response)
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

	// get params
	urlParams := r.URL.Query()
	if fontPath := urlParams.Get("font"); fontPath != "" {
		config.fontPath = fontPath
	}

	if scale := urlParams.Get("scale"); scale != "" {
		if scaleVal, err := strconv.ParseUint(scale, 10, 32); err == nil {
			config.scale = uint(scaleVal)
		}
	}

	if colored := urlParams.Get("colored"); colored != "" {
		config.colored = colored == "true"
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
