package convert

import (
	"image"
	"image/color"
)

type ConversionStrategy interface {
	convert(pixels [][]color.Color) image.Image
}

type ConversionContext struct {
	strategy ConversionStrategy
}

func (cc *ConversionContext) SetStrategy(cs ConversionStrategy) {
	cc.strategy = cs
}

func (cc *ConversionContext) ExecuteConversion(pixels [][]color.Color) image.Image {
	return cc.strategy.convert(pixels)
}

// Convert the image into a tensor to further manipulation
func ConvertIntoTensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	pixels := make([][]color.Color, size.Y)

	for y := 0; y < size.Y; y++ {
		pixels[y] = make([]color.Color, size.X)
		for x := 0; x < size.X; x++ {
			pixels[y][x] = img.At(x, y)
		}
	}

	return pixels
}
