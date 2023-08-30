package models

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type Rgba64Strategy struct{}

func (pstr *Rgba64Strategy) ConvertToModel(img image.Image) [][]color.Color {
	rect := img.Bounds()
	res := make([][]color.Color, rect.Max.Y)

	for y := 0; y < rect.Max.Y; y++ {
		res[y] = make([]color.Color, rect.Max.X)

		for x := 0; x < rect.Max.X; x++ {
			p := img.At(x, y)
			original, ok := color.RGBA64Model.Convert(p).(color.RGBA64)
			if ok {
				res[y][x] = original
			}
		}
	}

	return res
}
