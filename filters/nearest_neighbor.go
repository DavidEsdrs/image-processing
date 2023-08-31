package filters

import "image/color"

type NearestNeighborFilter struct {
	factor float32
}

func NewNearestNeighborFilter(factor float32) NearestNeighborFilter {
	return NearestNeighborFilter{}
}

func (nn NearestNeighborFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor

	proportion := int(1 / nn.factor)
	rows := int(len(img) / proportion)
	cols := int(len(img[0]) / proportion)

	res := make([][]color.Color, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]color.Color, cols)

		for j := 0; j < cols; j++ {
			res[i][j] = img[i*proportion][j*proportion]
		}
	}

	*tensor = res
	return nil
}
