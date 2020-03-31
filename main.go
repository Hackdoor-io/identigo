package main

import (
	"fmt"
	"hash/fnv"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

type RGBColor = struct {
	R float64
	G float64
	B float64
}

func stringToUint32(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func generateRGBFromString(str string, steps int) RGBColor {

	intList := stringToUint32(str)

	r := intList >> (2 * steps) & 0xFF
	g := intList >> (4 * steps) & 0xFF
	b := intList >> (6 * steps) & 0xFF

	return RGBColor{
		R: float64(r),
		G: float64(g),
		B: float64(b),
	}
}

func linearGradient(x, width int, initialColor, finalColor RGBColor) (uint8, uint8, uint8) {
	d := float64(x) / float64(width)

	r := initialColor.R + (d * (finalColor.R - initialColor.R))
	g := initialColor.G + (d * (finalColor.G - finalColor.G))
	b := initialColor.B + (d * (finalColor.B - initialColor.B))

	return uint8(r), uint8(g), uint8(b)
}

func generateImage(width int, height int, initialColor RGBColor, finalColor RGBColor) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 1; x <= width; x++ {
		for y := 1; y <= height; y++ {
			r, g, b := linearGradient(x, width, initialColor, finalColor)

			c := color.RGBA{r, g, b, 255}

			img.Set(x, y, c)
		}
	}

	return img
}

func main() {
	start := time.Now()
	str := "my_username"

	initialColor := generateRGBFromString(str, 2)
	finalColor := generateRGBFromString(str, 4)

	resultImage := generateImage(500, 500, initialColor, finalColor)

	img, _ := os.Create("new.png")
	defer img.Close()

	png.Encode(img, resultImage)
	duration := time.Since(start)
	fmt.Println(duration)

	fmt.Println(generateRGBFromString("4e3ef92ed95e0776ff69797b475ccd58", 1))
}
