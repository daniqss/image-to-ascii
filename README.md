# image to ascii
tool to generate ascii art images from other images
![example](./assets/image.png_ascii.png)

## how to run
```bash
go run . --scale 2 --print --colored path/to/image.jpg
# or
make run ARGS="--scale 2 --print --colored image.png path/to/image.jpg"
```

## endpoints and example on how to get from server
```bash
curl -X GET "http://localhost:3000/api/v1/fonts"

curl -X POST -F "image=@test/image.png" "http://localhost:3000/api/v1/?scale=16&colored=true&font=UbuntuNerdFont-Bold" --output output_ascii.png
```
there are 3 query params:
- scale: int
- colored: bool
- font: string

`font` must be one of the fonts that are available on the server, you can get the list of available fonts by calling the first endpoint