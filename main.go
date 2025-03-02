package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	mode     string
	path     string
	fontPath string
	scale    uint
	density  string
	print    bool
	colored  bool

	help bool
}

func main() {
	config, err := manageArgs(os.Args[1:])
	if config.help {
		help()
		return
	}
	if err != nil {
		help()
		log.Fatal(err)
	}

	switch config.mode {
	case "cli":
		useCliMode(config)
	case "server":
		useServerMode(config)
	default:
		{
			help()
			log.Fatal("invalid mode")
		}
	}
}

func manageArgs(args []string) (Config, error) {
	var config Config

	if len(args) == 0 {
		return config, errors.New("no arguments provided")
	}

	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	fs.StringVar(&config.mode, "mode", "cli", "Specify the mode (optional, default: cli)")
	fs.StringVar(&config.fontPath, "fontPath", "", "Wanted ttf font path (optional)")
	fs.UintVar(&config.scale, "scale", 8, "Specify the processing scale (optional, default: 8)")
	fs.StringVar(&config.density, "density", " .;coPO#@", "Specify the density (optional, default: \" .;coPO#@\")")
	fs.BoolVar(&config.print, "print", false, "Print the result (optional, default: false)")
	fs.BoolVar(&config.colored, "colored", false, "Enable colored output (optional, default: false)")
	fs.BoolVar(&config.help, "help", false, "Show this help message and exit")
	fs.BoolVar(&config.help, "h", false, "Show this help message and exit")

	config.density = config.density + " "

	// Parse flags
	err := fs.Parse(args)
	if err != nil {
		return config, err
	}

	// Ensure the last argument is treated as the path
	if config.mode != "server" {
		remainingArgs := fs.Args()
		if len(remainingArgs) == 0 {
			return config, errors.New("path is required")
		}
		config.path = remainingArgs[len(remainingArgs)-1]
		return config, nil
	}

	// If is selected the server mode, the path is not required
	config.path = ""
	return config, nil
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("  image-to-ascii [OPTIONS] <PATH>")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --mode string        Specify the mode (optional, default: cli)")
	fmt.Println("  --fontPath string    Wanted font path (optional, default: /usr/share/fonts/OpenSans-BoldItalic.ttf)")
	fmt.Println("  --scale uint8        Specify the processing scale (optional, default: 8)")
	fmt.Println("  --print              Print the result (optional, default: false)")
	fmt.Println("  --colored            Enable colored output (optional, default: false)")
	fmt.Println("  -h, --help          Show this help message and exit")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  image-to-ascii --scale 2 --print --colored image.png")
	fmt.Println("  image-to-ascii image.png")
	fmt.Println()
}
