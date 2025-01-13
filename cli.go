package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func useCliMode(config Config) {
	img, _, err := getImageFromPath(config.path)
	if err != nil {
		log.Fatal(err)
	}

	ascii := Ascii{
		config: config,
		img:    img,
	}

	if !ascii.config.print {
		if err := ascii.generateAscii(); err != nil {
			log.Fatal(err)
		}
	} else {
		ascii.printAscii()
	}
}

func getImageFromPath(filepath string) (image.Image, string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %w", err)
	}
	return img, format, err
}
