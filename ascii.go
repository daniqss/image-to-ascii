package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"unicode/utf8"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type Ascii struct {
	img image.Image
	config Config
}

func (ascii Ascii) generateAscii() error {
	width, height := ascii.img.Bounds().Dx(), ascii.img.Bounds().Dy()
	scaledW, scaledH := uint(width)/uint(ascii.config.scale), uint(height)/uint(ascii.config.scale)

	dc := gg.NewContext(width, height)
	dc.SetColor(color.Black)

	if err := dc.LoadFontFace(ascii.config.fontPath, float64(ascii.config.scale)); err != nil {
		return err
	}


	imgResized := resize.Resize(scaledW, scaledH, ascii.img, resize.Bilinear)
	for x := range scaledH {
		for y := range scaledW {
			c := imgResized.At(int(x), int(y))
			b := getBrightness(c)
			char := getCharFromBrightness(b)
			str := append(make([]byte, 1), char)

			dc.DrawString(string(str), float64(x * ascii.config.scale), float64(y * ascii.config.scale))
			
		}
		
	}

    dc.SavePNG(ascii.config.path + "ascii.png")
	return nil
}

func printAscii(img image.Image, config Config) {
	width, height := uint(img.Bounds().Max.X)/uint(config.scale), uint(img.Bounds().Max.Y)/uint(config.scale)
	imgResized := resize.Resize(width, height, img, resize.Bilinear)
	for x := range height {
		for y := range width {
			color := imgResized.At(int(y), int(x))
			
			if config.colored {
				printColoredBackground(color)
			} else {
				printAsciiChar(color)
			}
		}
		fmt.Printf("\033[0m\n")
	}

	fmt.Printf("\033[0m")
}

func printColoredBackground(color color.Color) {
	r, g, b, _ := color.RGBA()

	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	fmt.Printf("\033[48;2;%d;%d;%dm  ", r8, g8, b8)
}

func printAsciiChar(c color.Color) {
	b := getBrightness(c)
	char := getCharFromBrightness(b)
	fmt.Printf("%c%c", char, char)
}

func getBrightness(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	r8 := float64(r >> 8)
	g8 := float64(g >> 8)
	b8 := float64(b >> 8)

	return 0.299*r8 + 0.587*g8 + 0.114*b8
}

func getCharFromBrightness(b float64) byte {
	i := int(b / 255.0 * float64(utf8.RuneCountInString(DENSITY)-1))
	return DENSITY[i]
}