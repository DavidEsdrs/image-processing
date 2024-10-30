package filters

import "github.com/DavidEsdrs/image-processing/quad"

type FlipYFilter struct{}

func NewFlipYFilter() FlipYFilter {
	return FlipYFilter{}
}

func (fyf FlipYFilter) Execute(q *quad.Quad) error {
	rows, cols := q.Rows, q.Cols
	for y := 0; y < rows/2; y++ {
		for x := 0; x < cols; x++ {
			// Pega o pixel na posição (x, y) e o pixel na posição espelhada (x, rows-y-1)
			topPixel := q.GetPixel(x, y)
			bottomPixel := q.GetPixel(x, rows-y-1)

			// Troca os pixels
			q.SetPixel(x, y, bottomPixel)
			q.SetPixel(x, rows-y-1, topPixel)
		}
	}
	return nil
}
