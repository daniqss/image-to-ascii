package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Ascii struct {
	Width      int
	Height     int
	Color      string
	Brightness float64
	Character  byte
}

func new(img image.Image) Ascii {
	return Ascii{
		Width:  img.Bounds().Max.X,
		Height: img.Bounds().Max.Y,
	}
}

func rgbToGrey(r, g, b uint8) float64 {
	rf64 := float64(r)
	gf64 := float64(g)
	bf64 := float64(b)
	return 0.2989*rf64 + 0.5870*gf64 + 0.1140*bf64
}

// imageToColorMap(img image.Image)
