APP_NAME = image-to-ascii
SRC_DIR = .
GO_FILES = main.go ascii.go cli.go

.PHONY: all run build clean test deps fmt vet help

all: build

build: $(GO_FILES)
	@go build -o $(APP_NAME) $(SRC_DIR)

run: $(GO_FILES)
	@go run $(SRC_DIR) $(ARGS)

clean:
	@rm -f $(APP_NAME) image.png_ascii.png

test:
	@go test ./...

help:
	@go run $(SRC_DIR) --help 
