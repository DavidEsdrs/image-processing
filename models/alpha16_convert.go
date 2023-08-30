package models

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type Alpha16Strategy struct{}

func (pstr *Alpha16Strategy) ConvertToModel(img image.Image) [][]color.Color {
	rect := img.Bounds()
	res := make([][]color.Color, rect.Max.Y)

	for y := 0; y < rect.Max.Y; y++ {
		res[y] = make([]color.Color, rect.Max.X)

		for x := 0; x < rect.Max.X; x++ {
			p := img.At(x, y)
			original, ok := color.Alpha16Model.Convert(p).(color.Alpha16)
			if ok {
				res[y][x] = original
			}
		}
	}

	return res
}
