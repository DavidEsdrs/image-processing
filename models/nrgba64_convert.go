package models

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type Nrgba64Strategy struct{}

func (pstr *Nrgba64Strategy) ConvertToModel(img image.Image) [][]color.Color {
	rect := img.Bounds()
	res := make([][]color.Color, rect.Max.Y)

	for y := 0; y < rect.Max.Y; y++ {
		res[y] = make([]color.Color, rect.Max.X)

		for x := 0; x < rect.Max.X; x++ {
			p := img.At(x, y)
			original, ok := color.NRGBA64Model.Convert(p).(color.NRGBA64)
			if ok {
				res[y][x] = original
			}
		}
	}

	return res
}
