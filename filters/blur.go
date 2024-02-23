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
	if sigma == 1 {
		sigma = float64(kernelSize) / 2
	}
	bf := BlurFilter{l: l, sigma: sigma, kernelSize: kernelSize, kernel: kernel}
	return bf, nil
}

func (bf BlurFilter) Execute(tensor *[][]color.Color) error {
	height := len(*tensor)
	width := len((*tensor)[0])

	copy := deepCopy(tensor)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := bf.getValuesForPixel(tensor, &copy, x, y)
			(*tensor)[y][x] = color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			}
		}
	}

	return nil
}

// produces a deep copy from src to dst
func deepCopy(src *[][]color.Color) (copy [][]color.Color) {
	copy = make([][]color.Color, len(*src))
	for i := range copy {
		copy[i] = append(copy[i], (*src)[i]...)
	}
	return
}

func (bf *BlurFilter) getValuesForPixel(
	tensor *[][]color.Color,
	copy *[][]color.Color,
	startX,
	startY int,
) (r, g, b, a uint8) {
	height := len(*tensor)
	width := len((*tensor)[0])

	var (
		rnew uint8
		gnew uint8
		bnew uint8
		anew uint8
	)

	sy := fixed(startY - (bf.kernelSize / 2))
	sx := fixed(startX - (bf.kernelSize / 2))

	endY := fixed(startY + (bf.kernelSize / 2))
	endX := fixed(startX + (bf.kernelSize / 2))

	for y := sy; y <= endY && y < height; y++ {
		for x := sx; x <= endX && x < width; x++ {
			r, g, b, a := (*copy)[y][x].RGBA()

			w := bf.kernel[y-sy][x-sx]

			convertedR := float64(r) * w
			convertedG := float64(g) * w
			convertedB := float64(b) * w
			convertedA := float64(a) * w

			rnew += uint8(uint32(convertedR) >> 8)
			gnew += uint8(uint32(convertedG) >> 8)
			bnew += uint8(uint32(convertedB) >> 8)
			anew += uint8(uint32(convertedA) >> 8)
		}
	}

	return rnew, gnew, bnew, anew
}

func fixed(x int) int {
	if x < 0 {
		return 0
	}
	return x
}
