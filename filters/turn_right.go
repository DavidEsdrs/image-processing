package filters

import (
	"github.com/DavidEsdrs/image-processing/quad"
)

type TurnRightFilter struct{}

func NewTurnRightFilter() TurnRightFilter {
	return TurnRightFilter{}
}

func (gsf TurnRightFilter) Execute(q *quad.Quad) error {
	originalRows := q.Rows
	originalCols := q.Cols

	rows := q.Rows

	res := q.Clone()

	for i := 0; i < originalRows; i++ {
		for j := 0; j < originalCols; j++ {
			res.SetPixel(i, rows-j-1, q.GetPixel(j, i))
		}
	}

	*q = *res
	return nil
}
