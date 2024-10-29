package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/quad"
)

type NearestNeighborFilter struct {
	factor              float64
	width               int
	height              int
	needCalculateFactor bool
}

func NewNearestNeighborFilter(factor float64, width, height int) NearestNeighborFilter {
	nn := NearestNeighborFilter{factor: factor, width: width, height: height}
	nn.needCalculateFactor = factor != 1
	return nn
}

func (nn NearestNeighborFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor

	originalWidth := len(img[0])
	originalHeight := len(img)

	if nn.needCalculateFactor {
		nn.width = int(float64(originalWidth) * nn.factor)
		nn.height = int(float64(originalHeight) * nn.factor)
	}

	res := make([][]color.Color, nn.height)

	x_ratio := (originalWidth<<16)/nn.width + 1
	y_ratio := (originalHeight<<16)/nn.height + 1

	var x2 int
	var y2 int

	for y := 0; y < nn.height; y++ {
		res[y] = make([]color.Color, nn.width)
		for x := 0; x < nn.width; x++ {
			x2 = (x * x_ratio) >> 16
			y2 = (y * y_ratio) >> 16
			res[y][x] = img[y2][x2]
		}
	}

	*tensor = res
	return nil
}

func (nn NearestNeighborFilter) ExecuteFilter(q *quad.Quad) error {
	originalWidth := q.Cols
	originalHeight := q.Rows

	if nn.needCalculateFactor {
		nn.width = int(float64(originalWidth) * nn.factor)
		nn.height = int(float64(originalHeight) * nn.factor)
	}

	res := q.Clone()

	x_ratio := (originalWidth<<16)/nn.width + 1
	y_ratio := (originalHeight<<16)/nn.height + 1

	var x2 int
	var y2 int

	for y := 0; y < nn.height; y++ {
		for x := 0; x < nn.width; x++ {
			x2 = (x * x_ratio) >> 16
			y2 = (y * y_ratio) >> 16
			res.SetPixel(x, y, q.GetPixel(x2, y2))
		}
	}

	*q = *res

	return nil
}
