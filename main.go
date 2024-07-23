package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

const DEFAULT_SCALE = 8
const DENSITY = " .;coPO#@ "

type Config struct {
	path     string
	fontPath string
	scale    uint
	print    bool
	colored  bool
	edges    bool
}

func main() {
	config, err := manageArgs(os.Args[1:])
	if err != nil {
		help()
		log.Fatal(err)
	}

	img, _, err := getImageFromPath(config.path)
	if err != nil {
		log.Fatal(err)
	}

	ascii := Ascii{
		config: config,
		img: img,
	}

	if err := ascii.generateAscii(); err != nil {
		log.Fatal(err)
	}
}

func manageArgs(args []string) (Config, error) {
	var config Config

	if len(args) == 0 {
		return config, errors.New("no arguments provided")
	}

	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	fs.StringVar(&config.fontPath, "fontPath", "/usr/share/fonts/OpenSans-BoldItalic.ttf", "Wanted font path (optional, default: /usr/share/fonts/OpenSans-BoldItalic.ttf)")
	fs.UintVar(&config.scale, "scale", DEFAULT_SCALE, "Specify the processing scale (optional, default: 8)")
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
	fmt.Println()
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
