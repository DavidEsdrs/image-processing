package palette

import (
	"image"
	"image/color"
)

func GetPalette(img image.Image) (color.Palette, int) {
	uniqueColors := make(map[color.Color]bool)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			uniqueColors[pixelColor] = true
		}
	}
	plt := color.Palette{}
	for k := range uniqueColors {
		plt = append(plt, k)
	}
	return plt, len(plt)
}
