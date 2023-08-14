package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type Alpha16Strategy struct{}

func (pstr *Alpha16Strategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewAlpha16(rect)
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
			original, ok := color.Alpha16Model.Convert(p).(color.Alpha16)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
