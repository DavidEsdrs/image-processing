package filters

import (
	"image/color"
	"math"

	"github.com/DavidEsdrs/image-processing/quad"
)

type BrightnessFilter struct {
	luminance int
}

func NewBrightnessFilter(luminance int) BrightnessFilter {
	return BrightnessFilter{luminance}
}

// reference for this algorithm: https://ie.nitk.ac.in/blog/2020/01/19/algorithms-for-adjusting-brightness-and-contrast-of-an-image/
func (bf BrightnessFilter) Execute(q *quad.Quad) error {
	q.Apply(func(pixel color.RGBA) color.RGBA {
		r, g, b, a := pixel.RGBA()

		rAsUint8 := r >> 8
		gAsUint8 := g >> 8
		bAsUint8 := b >> 8
		A := uint8(a >> 8)

		newR := truncate(int(rAsUint8), bf.luminance)
		newG := truncate(int(gAsUint8), bf.luminance)
		newB := truncate(int(bAsUint8), bf.luminance)

		newPixel := color.RGBA{
			R: uint8(newR),
			G: uint8(newG),
			B: uint8(newB),
			A: A,
		}

		return newPixel
	})

	return nil
}

func truncate(x, y int) uint8 {
	result := x + y
	if result > math.MaxUint8 {
		return math.MaxUint8
	} else if result < 0 {
		return 0
	}
	return uint8(result)
}
