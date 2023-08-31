package filters

import "image/color"

type FlipXFilter struct{}

func NewFlipXFilter() FlipXFilter {
	return FlipXFilter{}
}

func (fyf FlipXFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor
	rows := len(img)
	cols := len(img[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols/2; j++ {
			img[i][cols-j-1], img[i][j] = img[i][j], img[i][cols-j-1]
		}
	}
	return nil
}
