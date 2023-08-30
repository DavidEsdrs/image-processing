package interfaces

import (
	"image"
	"image/color"
)

type Converter interface {
	ConvertToModel(img image.Image) [][]color.Color
}
