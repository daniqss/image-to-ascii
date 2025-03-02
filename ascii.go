package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"unicode/utf8"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type Ascii struct {
	img    image.Image
	config Config
}

func (ascii Ascii) generateAscii(w *io.Writer) error {
	width, height := ascii.img.Bounds().Dx(), ascii.img.Bounds().Dy()
	scaledW, scaledH := uint(width)/uint(ascii.config.scale), uint(height)/uint(ascii.config.scale)

	dc := gg.NewContext(width, height)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Default color if not using colored mode
	dc.SetColor(color.RGBA{R: 201, G: 91, B: 201, A: 255})

	// manage font in cli mode
	if ascii.config.mode == "cli" && ascii.config.fontPath != "" {
		if err := dc.LoadFontFace(ascii.config.fontPath, float64(ascii.config.scale)); err != nil {
			return err
		}
	}
	if ascii.config.mode == "server" {
		if err := dc.LoadFontFace("./fonts/"+ascii.config.fontPath+".ttf", float64(ascii.config.scale)); err != nil {
			return err
		}
	}

	imgResized := resize.Resize(scaledW, scaledH, ascii.img, resize.Bilinear)
	for y := range scaledH {
		for x := range scaledW {
			c := imgResized.At(int(x), int(y))

			// If colored mode is enabled, set the drawing color to match the pixel
			if ascii.config.colored {
				r, g, b, _ := c.RGBA()
				// Convert from 0-65535 range to 0-255 range
				dc.SetColor(color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: 255,
				})
			}

			b := ascii.getBrightness(c)
			char := ascii.getCharFromBrightness(b)
			str := append(make([]byte, 1), char)

			dc.DrawString(string(str), float64(x*ascii.config.scale), float64(y*ascii.config.scale))
		}
	}

	if ascii.config.mode == "server" {
		if w == nil {
			return fmt.Errorf("no writer provided")
		}

		err := dc.EncodePNG(*w)
		if err != nil {
			return err
		}
	} else {
		dc.SavePNG(ascii.config.path + "_ascii.png")
	}

	return nil
}

func (ascii Ascii) printAscii() {
	width, height := uint(ascii.img.Bounds().Max.X)/uint(ascii.config.scale), uint(ascii.img.Bounds().Max.Y)/uint(ascii.config.scale)
	imgResized := resize.Resize(width, height, ascii.img, resize.Bilinear)
	for x := range height {
		for y := range width {
			c := imgResized.At(int(y), int(x))

			if ascii.config.colored {
				fmt.Printf("%s", ascii.sprintColoredBackground(c))
			} else {
				ascii.printAsciiChar(c)
			}
		}
		fmt.Printf("\033[0m\n")
	}

	fmt.Printf("\033[0m")
}

func (ascii Ascii) sprintColoredBackground(color color.Color) string {
	r, g, b, _ := color.RGBA()

	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	return fmt.Sprintf("\033[48;2;%d;%d;%dm  ", r8, g8, b8)
}

func (ascii Ascii) printAsciiChar(c color.Color) {
	b := ascii.getBrightness(c)
	char := ascii.getCharFromBrightness(b)
	fmt.Printf("%c%c", char, char)
}

func (ascii Ascii) getBrightness(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	r8 := float64(r >> 8)
	g8 := float64(g >> 8)
	b8 := float64(b >> 8)

	return 0.299*r8 + 0.587*g8 + 0.114*b8
}

func (ascii Ascii) getCharFromBrightness(b float64) byte {
	i := int(b / 255.0 * float64(utf8.RuneCountInString(ascii.config.density)-1))
	return ascii.config.density[i]
}
