package filters

import (
	"image/color"
	"math"
)

type RotateFilter struct {
	degrees float64
}

func NewRotateFilter(degrees float64) RotateFilter {
	return RotateFilter{degrees}
}

func (rf RotateFilter) Execute(tensor *[][]color.Color) error {
	centerX, centerY := getCenterCoordinates(tensor)
	height := len(*tensor)
	width := len((*tensor)[0])

	var tensorCopy *[][]color.Color

	imgD := withDimensions(width, height)
	tensorCopy = &imgD

	for y1 := 0; y1 < height; y1++ {
		for x1 := 0; x1 < width; x1++ {
			x2, y2 := rotatePixel(rf.degrees, centerX, centerY, x1, y1)

			if x2 < width && y2 < height && x2 >= 0 && y2 >= 0 {
				(*tensorCopy)[y1][x1] = (*tensor)[y2][x2]
			}
		}
	}

	*tensor = *tensorCopy

	return nil
}

func withDimensions(width, height int) [][]color.Color {
	copy := make([][]color.Color, height)

	for y := 0; y < height; y++ {
		copy[y] = make([]color.Color, width)
	}

	return copy
}

func getCenterCoordinates(tensor *[][]color.Color) (x, y int) {
	height := len(*tensor)
	width := len((*tensor)[0])
	y = height / 2
	x = width / 2
	return
}

// x2 = x0 + cos(theta)*(x1-x0) + sin(theta)*(y1-y0)
// y2 = y0 - sin(theta)*(x1-x0) + cos(theta)*(y1-y0)
func rotatePixel(degrees float64, centerX, centerY, x1, y1 int) (x2, y2 int) {
	rads := toRadians(degrees)

	x2 = int(math.Cos(rads)*float64(x1-centerX) + float64(y1-centerY)*math.Sin(rads))
	y2 = int(math.Sin(rads)*-float64(x1-centerX) + float64(y1-centerY)*math.Cos(rads))

	x2 += centerX
	y2 += centerY

	return
}

func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
