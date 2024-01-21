package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/utils"
)

type BlurFilter struct {
	l          logger.Logger
	sigma      float64
	kernelSize int // this will define how blurry the image is
	kernel     [][]float64
}

func NewBlurFilter(l logger.Logger, sigma float64, kernelSize int) (BlurFilter, error) {
	kernel := utils.GaussianKernel(kernelSize, sigma)
	bf := BlurFilter{l: l, sigma: sigma, kernelSize: kernelSize, kernel: kernel}
	return bf, nil
}

func (bf BlurFilter) Execute(tensor *[][]color.Color) error {
	height := len(*tensor)
	width := len((*tensor)[0])

	copy := make([][]color.Color, height)
	for i := range copy {
		copy[i] = append(copy[i], (*tensor)[i]...)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, _, _, a := (*tensor)[y][x].RGBA()
			r, g, b := bf.getValuesForPixel(tensor, &copy, x, y)
			(*tensor)[y][x] = color.RGBA{
				R: r,
				G: g,
				B: b,
				A: uint8(a >> 8),
			}
		}
	}

	return nil
}

func (bf *BlurFilter) getValuesForPixel(
	tensor *[][]color.Color,
	copy *[][]color.Color,
	startX,
	startY int,
) (r, g, b uint8) {
	height := len(*tensor)
	width := len((*tensor)[0])

	var (
		rnew uint8
		gnew uint8
		bnew uint8
	)

	for y := startY; y < startY+bf.kernelSize && y < height; y++ {
		for x := startX; x < startX+bf.kernelSize && x < width; x++ {
			r, g, b, _ := (*copy)[y][x].RGBA()

			rc := uint8(r >> 8)
			gc := uint8(g >> 8)
			bc := uint8(b >> 8)

			w := bf.kernel[y-startY][x-startX]

			convertedR := float64(rc) * w
			convertedG := float64(gc) * w
			convertedB := float64(bc) * w

			rnew += uint8(convertedR)
			gnew += uint8(convertedG)
			bnew += uint8(convertedB)
		}
	}

	return rnew, gnew, bnew
}
