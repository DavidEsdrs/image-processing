package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/DavidEsdrs/image-processing/quad"
)

type Alpha16Strategy struct{}

func (pstr *Alpha16Strategy) Convert(pixels *quad.Quad) image.Image {
	rect := image.Rect(0, 0, pixels.Cols, pixels.Rows)
	nImg := image.NewAlpha16(rect)
	for y := 0; y < pixels.Rows; y++ {
		for x := 0; x < pixels.Cols; x++ {
			p := pixels.GetPixel(x, y)
			original, ok := color.Alpha16Model.Convert(p).(color.Alpha16)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
