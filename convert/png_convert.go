package convert

import (
	"image"
	"image/color"
	_ "image/png"
)

type PngStrategy struct{}

// let's assume that png images are 48 bits depth
// Change it!!!
func (pstr *PngStrategy) convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewNRGBA(rect)
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
			original, ok := color.NRGBAModel.Convert(p).(color.NRGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
