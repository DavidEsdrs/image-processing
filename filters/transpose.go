package filters

import "image/color"

type TransposeFilter struct{}

func NewTransposeFilter() TransposeFilter {
	return TransposeFilter{}
}

func (gsf TransposeFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor

	res := make([][]color.Color, len(img[0]))

	for i := range res {
		res[i] = make([]color.Color, len(img))
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			res[j][i] = img[i][j]
		}
	}

	*tensor = res
	return nil
}
