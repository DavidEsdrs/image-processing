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

	indexMap := make(map[color.Color]uint8)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			i := nImg.PixOffset(x, y)
			pixel := pixels[y][x]

			index, ok := indexMap[pixel]

			if !ok {
				index = uint8(pstr.palette.Index(pixel))
				indexMap[pixel] = index
			}

			nImg.Pix[i] = index
		}
	}

	return nImg
}
