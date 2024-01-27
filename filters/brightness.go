package filters

import (
	"image/color"
	"math"
)

type BrightnessFilter struct {
	luminance float64
}

func NewBrightnessFilter(luminance float64) BrightnessFilter {
	return BrightnessFilter{luminance}
}

// reference for this algorithm: https://ie.nitk.ac.in/blog/2020/01/19/algorithms-for-adjusting-brightness-and-contrast-of-an-image/
func (bf BrightnessFilter) Execute(tensor *[][]color.Color) error {
	height := len(*tensor)
	width := len((*tensor)[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := (*tensor)[y][x]
			r, g, b, a := pixel.RGBA()

			rAsUint8 := r >> 8
			gAsUint8 := g >> 8
			bAsUint8 := b >> 8
			A := uint8(a >> 8)

			newR := truncate(float64(rAsUint8), bf.luminance)
			newG := truncate(float64(gAsUint8), bf.luminance)
			newB := truncate(float64(bAsUint8), bf.luminance)

			newPixel := color.RGBA{
				R: uint8(newR),
				G: uint8(newG),
				B: uint8(newB),
				A: A,
			}
			(*tensor)[y][x] = newPixel
		}
	}

	return nil
}

func truncate(x, y float64) uint8 {
	result := x + y
	if result > math.MaxUint8 {
		return math.MaxUint8
	} else if result < 0 {
		return 0
	}
	return uint8(result)
}
