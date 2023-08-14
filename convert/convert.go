package convert

import (
	"fmt"
	"image"
	"image/color"
)

type ConversionStrategy interface {
	Convert(pixels [][]color.Color) image.Image
}

type ConversionContext struct{}

func NewConversionContext() *ConversionContext {
	return &ConversionContext{}
}

func (cc *ConversionContext) GetConversor(img image.Image) (ConversionStrategy, error) {
	model := img.ColorModel()
	switch model {
	case color.Alpha16Model:
		return &Alpha16Strategy{}, nil
	case color.AlphaModel:
		return &AlphaStrategy{}, nil
	case color.CMYKModel:
		return &CmykStrategy{}, nil
	case color.Gray16Model:
		return &Gray16Strategy{}, nil
	case color.GrayModel:
		return &GrayStrategy{}, nil
	case color.NRGBA64Model:
		return &Nrgba64Strategy{}, nil
	case color.NRGBAModel:
		return &NrgbaStrategy{}, nil
	case color.RGBA64Model:
		return &Rgba64Strategy{}, nil
	case color.RGBAModel:
		return &RgbaStrategy{}, nil
	case color.YCbCrModel:
		return &YcbcrStrategy{}, nil
	case color.NYCbCrAModel:
		return nil, fmt.Errorf("unsupported color model")
	}
	return nil, fmt.Errorf("unsupported color model")
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
