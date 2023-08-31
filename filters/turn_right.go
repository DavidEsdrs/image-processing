package filters

import "image/color"

type TurnLeftFilter struct{}

func NewTurnLeftFilter() TurnLeftFilter {
	return TurnLeftFilter{}
}

func (gsf TurnLeftFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor

	originalRows := len(img)
	originalCols := len(img[0])

	rows := len(img[0])
	cols := len(img)

	res := make([][]color.Color, rows)

	for i := range res {
		res[i] = make([]color.Color, cols)
	}

	for i := 0; i < originalRows; i++ {
		for j := 0; j < originalCols; j++ {
			res[j][originalRows-i-1] = img[i][j]
		}
	}

	*tensor = res
	return nil
}
