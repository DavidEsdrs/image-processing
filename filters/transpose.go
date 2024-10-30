package filters

import (
	"github.com/DavidEsdrs/image-processing/quad"
)

type TransposeFilter struct{}

func NewTransposeFilter() TransposeFilter {
	return TransposeFilter{}
}

func (gsf TransposeFilter) Execute(q *quad.Quad) error {
	res := q.Clone()

	for i := 0; i < q.Rows; i++ {
		for j := 0; j < q.Cols; j++ {
			res.SetPixel(i, j, q.GetPixel(j, i))
		}
	}

	*q = *res
	return nil
}
