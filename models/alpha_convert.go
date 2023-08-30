package models

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type AlphaStrategy struct{}

func (pstr *AlphaStrategy) ConvertToModel(img image.Image) [][]color.Color {
	rect := img.Bounds()
	nImg := image.NewAlpha(rect)
	res := make([][]color.Color, rect.Max.Y)

	for y := 0; y < rect.Max.Y; y++ {
		res[y] = make([]color.Color, rect.Max.X)

		for x := 0; x < rect.Max.X; x++ {
			p := nImg.At(x, y)
			original, ok := color.AlphaModel.Convert(p).(color.Alpha)
			if ok {
				res[y][x] = original
			}
		}
	}

	return res
}
