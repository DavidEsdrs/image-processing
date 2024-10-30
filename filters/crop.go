package filters

import (
	"fmt"
	"image"

	"github.com/DavidEsdrs/image-processing/quad"
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

func (cf CropFilter) Execute(img *quad.Quad) error {
	croppedWidth := cf.xend - cf.xstart
	croppedHeight := cf.yend - cf.ystart

	if cf.xstart < 0 || cf.ystart < 0 || cf.xend > img.Cols || cf.yend > img.Rows {
		return fmt.Errorf("crop points are out of image bounds")
	}

	croppedQuad := quad.NewQuad(croppedWidth, croppedHeight)

	for y := cf.ystart; y < cf.yend; y++ {
		for x := cf.xstart; x < cf.xend; x++ {
			pixel := img.GetPixel(x, y)
			croppedQuad.SetPixel(x-cf.xstart, y-cf.ystart, pixel)
		}
	}

	*img = *croppedQuad
	return nil
}
