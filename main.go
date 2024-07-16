package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"unicode/utf8"

	"github.com/nfnt/resize"
)

const DEFAULT_SCALE = 8
const DENSITY = " .;coPO?@■"

type Config struct {
	path    string
	scale   uint
	print   bool
	colored bool
	edges   bool
}

func main() {
	config, err := manageArgs(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
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

	if len(args) == 0 {
		return config, errors.New("no arguments provided")
	}

	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	fs.UintVar(&config.scale, "scale", 8, "Specify the processing scale (optional, default: 8)")
	fs.BoolVar(&config.print, "print", false, "Print the result (optional, default: false)")
	fs.BoolVar(&config.colored, "colored", false, "Enable colored output (optional, default: false)")
	fs.BoolVar(&config.edges, "edges", false, "Show only the edges (optional, default: false)")

	// Parse flags
	err := fs.Parse(args)
	if err != nil {
		return config, err
	}

	// Ensure the last argument is treated as the path
	remainingArgs := fs.Args()
	if len(remainingArgs) == 0 {
		return config, errors.New("path is required")
	}
	config.path = remainingArgs[len(remainingArgs)-1]

	return config, nil
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("  image-to-ascii [OPTIONS] <PATH>")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -h, --help          Show this help message and exit")
	fmt.Println("  -scale uint8        Specify the processing scale (optional, default: 8)")
	fmt.Println("  -print              Print the result (optional, default: false)")
	fmt.Println("  -colored            Enable colored output (optional, default: false)")
	fmt.Println("  -edges              Show only the edges (optional, default: false)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  image-to-ascii --scale 2 --print --colored image.png")
	fmt.Println("  image-to-ascii --edges image.png")
}
