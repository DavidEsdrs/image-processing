package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type GrayStrategy struct{}

func (pstr *GrayStrategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewGray(rect)
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
			original, ok := color.GrayModel.Convert(p).(color.Gray)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
