package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/hsl"
)

type SaturationFilter struct {
	saturation int
}

func NewSaturationFilter(saturation int) SaturationFilter {
	return SaturationFilter{saturation}
}

func (bf SaturationFilter) Execute(tensor *[][]color.Color) error {
	height := len(*tensor)
	width := len((*tensor)[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := (*tensor)[y][x]
			hslPixel := hsl.ColorToHsl(pixel)
			(*tensor)[y][x] = hslPixel
		}
	}

	return nil
}
