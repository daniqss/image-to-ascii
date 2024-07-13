package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/nfnt/resize"
)

const DEFAULT_SCALE = 8
const DENSITY = " .;coPO?@■"

type Config struct {
	path    string
	scale   uint8
	print   bool
	colored bool
	edges   bool
}

func main() {
	config, err := manageArgs(os.Args[1:])
	if err != nil {
		help()
		return
	}

	img, _, err := getImageFromPath(config.path)
	if err != nil {
		fmt.Printf("Error decoding image: %s\n", err)
		os.Exit(1)
	}

	printAscii(img, config)
}

func printAscii(img image.Image, config Config) {
	width, height := uint(img.Bounds().Max.X)/uint(config.scale), uint(img.Bounds().Max.Y)/uint(config.scale)
	imgResized := resize.Resize(width, height, img, resize.Bilinear)
	for x := range height {
		for y := range width {
			color := imgResized.At(int(y), int(x))
			// printColoredBackground(color, int(y), int(x))
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
	r, g, b, _ := c.RGBA()

	r8 := float64(r >> 8)
	g8 := float64(g >> 8)
	b8 := float64(b >> 8)

	brightness := 0.299*r8 + 0.587*g8 + 0.114*b8
	i := int(brightness / 255.0 * float64(utf8.RuneCountInString(DENSITY)-1))
	asciiChar := DENSITY[i]

	fmt.Printf("%c%c", asciiChar, asciiChar)
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

func manageArgs(args []string) (Config, error) {
	var config Config
	length := len(args)

	for i := 0; i < length-1; i++ {
		switch args[i] {
		case "--scale", "-s":
			if i <= length-1 {
				scale, err := strconv.Atoi(args[i+1])
				if err != nil {
					config.scale = DEFAULT_SCALE
				} else {
					config.scale = uint8(scale)
				}
			}
		case "--print", "-p":
			config.print = true

		case "--colored", "-c":
			config.colored = true

		case "--edges", "-e":
			config.edges = true

		}
	}

	return config, nil
}

func help() {
	fmt.Println("Usage: go run . <image-file>")
}
