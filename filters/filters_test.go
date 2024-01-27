package filters_test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/DavidEsdrs/image-processing/filters"
)

func TestHSL(t *testing.T) {
	red := color.RGBA{24, 98, 118, 255}
	redAsHsl := filters.ColorToHsl(red)
	r, g, b, a := redAsHsl.RGBA()
	res := color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
	fmt.Printf("hsl: %#v\n", redAsHsl)
	fmt.Printf("rgba: RGBA{%v, %v, %v, %v}\n", res.R, res.G, res.B, res.A)
}
