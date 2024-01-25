package filters

import "image/color"

type BrightnessFilter struct {
	luminance uint8
}

func NewBrightnessFilter(luminance uint8) BrightnessFilter {
	return BrightnessFilter{luminance}
}

func (bf BrightnessFilter) Execute(tensor *[][]color.Color) error {
	height := len(*tensor)
	width := len((*tensor)[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := (*tensor)[y][x]
			pixelAsYCbCr, ok := color.YCbCrModel.Convert(pixel).(color.YCbCr)
			if ok {
				newPixel := color.YCbCr{
					Y:  pixelAsYCbCr.Y * bf.luminance,
					Cb: pixelAsYCbCr.Cb,
					Cr: pixelAsYCbCr.Cr,
				}
				(*tensor)[y][x] = newPixel
			}
		}
	}
	return nil
}
