package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type CmykStrategy struct{}

func (pstr *CmykStrategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewCMYK(rect)
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
			original, ok := color.CMYKModel.Convert(p).(color.CMYK)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
