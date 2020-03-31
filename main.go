package main

import (
	"hash/fnv"
	"image"
	"image/color"
)

// RGBColor - an RGB color struct
type RGBColor = struct {
	R float64
	G float64
	B float64
}

// stringToUint32
// takes a string as input and returns an uint32 representation
func stringToUint32(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// generateRGBFromString
// generates an RGB code (using the RGBColor struct) given a string
// and an int (steps)
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

// linearGradient
// generates a single pixel for the final color gradient
func linearGradient(x, width int, initialColor, finalColor RGBColor) (uint8, uint8, uint8) {
	d := float64(x) / float64(width)

	r := initialColor.R + (d * (finalColor.R - initialColor.R))
	g := initialColor.G + (d * (finalColor.G - finalColor.G))
	b := initialColor.B + (d * (finalColor.B - initialColor.B))

	return uint8(r), uint8(g), uint8(b)
}

// generateImage
// creates the image pixel by pixel
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

// GenerateFromString - generates a new identicon given a string, width, and height
func GenerateFromString(str string, width int, height int) *image.RGBA {
	initialColor := generateRGBFromString(str, 2)
	finalColor := generateRGBFromString(str, 4)

	return generateImage(500, 500, initialColor, finalColor)
}
