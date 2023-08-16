package processor

import (
	"image/color"
)

type Processor interface {
	// This function is util when one of the filters changes the color model
	// So further - when getting the correct conversor - we can use this function
	// and get the correct color model
	GetColorModel() color.Model
	SetColorModel(colorModel color.Model)
	Crop(xstart, xend, ystart, yend int)
	FlipX()
	FlipY()
	TurnLeft()
	TurnRight()
	Transpose()
	Grayscale16()
	NearestNeighbor(factor float32)
	Execute(source *[][]color.Color) [][]color.Color
}

type Process func(*[][]color.Color)
