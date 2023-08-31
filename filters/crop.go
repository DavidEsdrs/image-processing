package filters

import (
	"fmt"
	"image"
	"image/color"
)

type CropFilter struct {
	xstart, ystart, xend, yend int
}

func NewCropFilter(inputImg image.Image, xstart, ystart, xend, yend int) (CropFilter, error) {
	rect := inputImg.Bounds()
	startPoint := image.Point{X: xstart, Y: ystart}
	endPoint := image.Point{X: xend, Y: yend}

	if !startPoint.In(rect) || !endPoint.In(rect) || endPoint.X == 0 || endPoint.Y == 0 {
		return CropFilter{}, fmt.Errorf("crop points are not in the image")
	}

	return CropFilter{xstart, ystart, xend, yend}, nil
}

func (cf CropFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor

	res := make([][]color.Color, cf.yend-cf.ystart)

	for i := range res {
		res[i] = make([]color.Color, cf.xend-cf.xstart)
	}

	for i := cf.ystart; i < cf.yend; i++ {
		for j := cf.xstart; j < cf.xend; j++ {
			res[i-cf.ystart][j-cf.xstart] = img[i][j]
		}
	}

	*tensor = res
	return nil
}
