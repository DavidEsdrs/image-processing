package filters

import "image/color"

type ProcessorWithColorModel interface {
	SetColorModel(colorModel color.Model)
}

func Grayscale16(pImg *[][]color.Color, ip ProcessorWithColorModel) {
	img := *pImg

	rows := int(len(img))
	cols := int(len(img[0]))

	res := make([][]color.Color, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]color.Color, cols)

		for j := 0; j < cols; j++ {
			originalColor := img[i][j]

			newColor := color.Gray16Model.Convert(originalColor)

			res[i][j] = newColor
		}
	}

	*pImg = res
	ip.SetColorModel(color.Gray16Model)
}

func Grayscale8(pImg *[][]color.Color, ip ProcessorWithColorModel) {
	img := *pImg

	rows := int(len(img))
	cols := int(len(img[0]))

	res := make([][]color.Color, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]color.Color, cols)

		for j := 0; j < cols; j++ {
			originalColor := img[i][j]

			newColor := color.GrayModel.Convert(originalColor)

			res[i][j] = newColor
		}
	}

	*pImg = res
	ip.SetColorModel(color.GrayModel)
}
