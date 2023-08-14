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
	rows := len(pixels)
	cols := len(pixels[0])
	rect := image.Rect(0, 0, cols, rows)
	nImg := image.NewPaletted(rect, pstr.palette)

	conversions := make(map[color.Color]color.Color, len(pstr.palette))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			p := pixels[y][x]
			c, ok := conversions[p]
			if !ok {
				// Adicionar uma conversão prévia se não existir
				c = pstr.palette.Convert(p)
				conversions[p] = c
			}
			nImg.Set(x, y, c)
		}
	}

	return nImg
}
