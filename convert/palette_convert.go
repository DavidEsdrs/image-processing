package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/DavidEsdrs/image-processing/quad"
)

type PaletteStrategy struct {
	palette color.Palette
}

func (pstr *PaletteStrategy) Convert(pixels *quad.Quad) image.Image {
	rect := image.Rect(0, 0, pixels.Cols, pixels.Rows)
	nImg := image.NewPaletted(rect, pstr.palette)

	indexMap := make(map[color.Color]uint8)

	for y := 0; y < pixels.Rows; y++ {
		for x := 0; x < pixels.Cols; x++ {
			i := nImg.PixOffset(x, y)
			pixel := pixels.GetPixel(x, y)

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
