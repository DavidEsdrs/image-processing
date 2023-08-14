package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type Rgba64Strategy struct{}

func (pstr *Rgba64Strategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewNRGBA64(rect)
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
			original, ok := color.RGBA64Model.Convert(p).(color.RGBA64)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
