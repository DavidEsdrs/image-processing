package filters

import "image/color"

type FlipYFilter struct{}

func NewFlipYFilter() FlipYFilter {
	return FlipYFilter{}
}

func (fyf FlipYFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor
	rows := len(img)
	cols := len(img[0])

	for i := 0; i < rows/2; i++ {
		for j := 0; j < cols; j++ {
			img[rows-i-1][j], img[i][j] = img[i][j], img[rows-i-1][j]
		}
	}
	return nil
}
