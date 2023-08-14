package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type YcbcrStrategy struct{}

func (pstr *YcbcrStrategy) Convert(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewYCbCr(rect, image.YCbCrSubsampleRatio444)
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[0]); x++ {
			p := pixels[y][x]
			ycbcrColor, ok := p.(color.YCbCr)
			if ok {
				offset := nImg.YOffset(x, y)
				nImg.Y[offset] = ycbcrColor.Y
				nImg.Cb[offset] = ycbcrColor.Cb
				nImg.Cr[offset] = ycbcrColor.Cr
			}
		}
	}
	return nImg
}
