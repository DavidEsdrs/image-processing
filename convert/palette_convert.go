package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type PaletteStrategy struct {
	palette color.Palette
}

func (pstr *PaletteStrategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewPaletted(rect, pstr.palette)
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[0]); x++ {
			q := pixels[y]
			if q == nil {
				continue
			}
			p := pixels[y][x]
			if p == nil {
				continue
			}
			clr := pstr.palette.Convert(p)
			nImg.Set(x, y, clr)
		}
	}
	return nImg
}
