package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/processor"
	"github.com/DavidEsdrs/image-processing/quad"
)

type Grayscale16Filter struct {
	context *processor.Invoker
}

func NewGrayscale16Filter(ctx *processor.Invoker) Grayscale16Filter {
	return Grayscale16Filter{ctx}
}

func (gs Grayscale16Filter) Execute(q *quad.Quad) error {
	var rows int
	var cols int

	rows = q.Rows
	cols = q.Cols

	res := q.Clone()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			originalColor := q.GetPixel(j, i)

			newColor := color.Gray16Model.Convert(originalColor)

			r, g, b, a := newColor.RGBA()

			res.SetPixel(j, i, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	*q = *res
	return nil
}
