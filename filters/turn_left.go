package filters

import (
	"github.com/DavidEsdrs/image-processing/quad"
)

type TurnLeftFilter struct{}

func NewTurnLeftFilter() TurnLeftFilter {
	return TurnLeftFilter{}
}

func (gsf TurnLeftFilter) Execute(q *quad.Quad) error {
	originalRows := q.Rows
	originalCols := q.Cols

	rows := q.Cols
	cols := q.Rows

	res := quad.NewQuad(cols, rows)

	for i := 0; i < originalRows; i++ {
		for j := 0; j < originalCols; j++ {
			originalPixel := q.GetPixel(j, i)
			res.SetPixel(originalRows-i-1, j, originalPixel)
		}
	}

	*q = *res
	return nil
}
