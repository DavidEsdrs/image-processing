package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/processor"
)

type Grayscale16Filter struct {
	context *processor.Invoker
}

func NewGrayscale16Filter(ctx *processor.Invoker) Grayscale16Filter {
	return Grayscale16Filter{ctx}
}

func (gs Grayscale16Filter) Execute(tensor *[][]color.Color) error {
	img := *tensor

	var rows int
	var cols int

	rows = int(len(img))
	cols = int(len(img[0]))

	res := make([][]color.Color, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]color.Color, cols)

		for j := 0; j < cols; j++ {
			originalColor := img[i][j]

			newColor := color.Gray16Model.Convert(originalColor)

			res[i][j] = newColor
		}
	}

	*tensor = res
	return nil
}
