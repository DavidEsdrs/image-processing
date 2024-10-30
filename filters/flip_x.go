package filters

import "github.com/DavidEsdrs/image-processing/quad"

type FlipXFilter struct{}

func NewFlipXFilter() FlipXFilter {
	return FlipXFilter{}
}

func (fyf FlipXFilter) Execute(q *quad.Quad) error {
	rows, cols := q.Rows, q.Cols
	for y := 0; y < rows; y++ {
		for x := 0; x < cols/2; x++ {
			// Pega o pixel da posição (x, y) e o pixel da posição espelhada (cols-x-1, y)
			leftPixel := q.GetPixel(x, y)
			rightPixel := q.GetPixel(cols-x-1, y)

			// Troca os pixels
			q.SetPixel(x, y, rightPixel)
			q.SetPixel(cols-x-1, y, leftPixel)
		}
	}
	return nil
}
