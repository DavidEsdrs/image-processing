package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/hsl"
	"github.com/DavidEsdrs/image-processing/quad"
	"github.com/DavidEsdrs/image-processing/utils"
)

// SaturationFilter adjusts the saturation level of an image.
type SaturationFilter struct {
	// Saturation value (-1.0 to 1.0)
	saturation float64
}

func NewSaturationFilter(saturation float64) SaturationFilter {
	return SaturationFilter{saturation}
}

func (bf SaturationFilter) Execute(tensor *[][]color.Color) error {
	height := len(*tensor)
	width := len((*tensor)[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := (*tensor)[y][x]
			hslPixel := hsl.ColorToHsl(pixel)
			hslPixel.S += bf.saturation
			hslPixel.S = utils.Clamp(hslPixel.S, 0.0, 1.0)
			(*tensor)[y][x] = hslPixel
		}
	}

	return nil
}

func (bf SaturationFilter) ExecutePixel(img *quad.Quad) error {
	img.Apply(func(pixel color.RGBA) color.RGBA {
		hslPixel := hsl.ColorToHsl(pixel)
		hslPixel.S += bf.saturation
		hslPixel.S = utils.Clamp(hslPixel.S, 0.0, 1.0)
		r, g, b, a := hslPixel.RGBA()
		return color.RGBA{
			R: uint8(r << 8),
			G: uint8(g << 8),
			B: uint8(b << 8),
			A: uint8(a << 8),
		}
	})
	return nil
}
