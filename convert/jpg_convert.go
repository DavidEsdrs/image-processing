package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
)

type JpgStrategy struct{}

// let's assume that jpg images are 24 bits depth
// Change it!!!
func (pstr *JpgStrategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewRGBA(rect)
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
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
