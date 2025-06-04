package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 800
	height := 200

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Define start and end colors (RGBA)
	startColor := color.RGBA{R: 73, G: 24, B: 108, A: 255} // Dark purple
	endColor := color.RGBA{R: 173, G: 216, B: 230, A: 255} // Light blue

	for x := 0; x < width; x++ {
		// Linear interpolation ratio
		t := float64(x) / float64(width-1)

		// Interpolate each color component
		r := uint8(float64(startColor.R)*(1-t) + float64(endColor.R)*t)
		g := uint8(float64(startColor.G)*(1-t) + float64(endColor.G)*t)
		b := uint8(float64(startColor.B)*(1-t) + float64(endColor.B)*t)
		a := uint8(float64(startColor.A)*(1-t) + float64(endColor.A)*t)

		// Set the color for each vertical line
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{r, g, b, a})
		}
	}

	// Save to file
	f, err := os.Create("gradient.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

	println("Gradient image saved to gradient.png")
}
