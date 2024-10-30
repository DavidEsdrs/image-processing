package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/DavidEsdrs/image-processing/quad"
)

type RgbaStrategy struct{}

func (pstr *RgbaStrategy) Convert(pixels *quad.Quad) image.Image {
	rect := image.Rect(0, 0, pixels.Cols, pixels.Rows)
	nImg := image.NewRGBA(rect)
	for y := 0; y < pixels.Rows; y++ {
		for x := 0; x < pixels.Cols; x++ {
			p := pixels.GetPixel(x, y)
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
