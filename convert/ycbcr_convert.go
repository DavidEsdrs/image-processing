package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/DavidEsdrs/image-processing/quad"
)

type YcbcrStrategy struct {
	subsamplingRatio image.YCbCrSubsampleRatio
}

func (pstr *YcbcrStrategy) Convert(pixels *quad.Quad) image.Image {
	rect := image.Rect(0, 0, pixels.Cols, pixels.Rows)
	nImg := image.NewYCbCr(rect, pstr.subsamplingRatio)

	for y := 0; y < pixels.Rows; y++ {
		for x := 0; x < pixels.Cols; x++ {
			p := pixels.GetPixel(x, y)
			ycbcrColor, ok := color.YCbCrModel.Convert(p).(color.YCbCr)
			if ok {
				yoffset := nImg.YOffset(x, y)
				coffset := nImg.COffset(x, y)
				nImg.Y[yoffset] = ycbcrColor.Y
				nImg.Cb[coffset] = ycbcrColor.Cb
				nImg.Cr[coffset] = ycbcrColor.Cr
			}
		}
	}
	return nImg
}
