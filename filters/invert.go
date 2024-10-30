package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/quad"
)

type InvertFilter struct{}

func NewInvertFilter() *InvertFilter {
	return &InvertFilter{}
}

func (i *InvertFilter) Execute(img *quad.Quad) error {
	img.Apply(func(pixel color.RGBA) color.RGBA {
		r, g, b, a := pixel.RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		a >>= 8
		r = 255 - r
		g = 255 - g
		b = 255 - b
		return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	})
	return nil
}
