package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/DavidEsdrs/image-processing/quad"
)

type Rgba64Strategy struct{}

func (pstr *Rgba64Strategy) Convert(pixels *quad.Quad) image.Image {
	rect := image.Rect(0, 0, pixels.Cols, pixels.Rows)
	nImg := image.NewRGBA64(rect)
	for y := 0; y < pixels.Rows; y++ {
		for x := 0; x < pixels.Cols; x++ {
			p := pixels.GetPixel(x, y)
			original, ok := color.RGBA64Model.Convert(p).(color.RGBA64)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
