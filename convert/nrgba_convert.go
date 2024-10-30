package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/DavidEsdrs/image-processing/quad"
)

type NrgbaStrategy struct{}

func (pstr *NrgbaStrategy) Convert(pixels *quad.Quad) image.Image {
	rect := image.Rect(0, 0, pixels.Cols, pixels.Rows)
	nImg := image.NewNRGBA(rect)
	for y := 0; y < pixels.Rows; y++ {
		for x := 0; x < pixels.Cols; x++ {
			p := pixels.GetPixel(x, y)
			original, ok := color.NRGBAModel.Convert(p).(color.NRGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
