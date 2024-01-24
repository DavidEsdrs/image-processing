package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type YcbcrStrategy struct {
	subsamplingRatio image.YCbCrSubsampleRatio
}

func (pstr *YcbcrStrategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewYCbCr(rect, pstr.subsamplingRatio)
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[0]); x++ {
			p := pixels[y][x]
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
