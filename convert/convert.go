package convert

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

type ConversionStrategy interface {
	Convert(pixels [][]color.Color) image.Image
}

type ConversionContext struct{}

func NewConversionContext() *ConversionContext {
	return &ConversionContext{}
}

func (cc *ConversionContext) GetConversor(file string) (ConversionStrategy, error) {
	strs := strings.Split(file, ".")
	ftype := strs[len(strs)-1]

	switch ftype {
	case "png":
		return &PngStrategy{}, nil
	case "jpeg":
	case "jpg":
		return &JpgStrategy{}, nil
	}

	return nil, fmt.Errorf("unknown file type")
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
