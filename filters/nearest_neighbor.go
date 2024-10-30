package filters

import (
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

func (nn NearestNeighborFilter) Execute(q *quad.Quad) error {
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
