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
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
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
) (r, g, b, a uint32) {
	height := len(*tensor)
	width := len((*tensor)[0])

	var (
		rnew uint32
		gnew uint32
		bnew uint32
		anew uint32
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

			rnew += uint32(convertedR)
			gnew += uint32(convertedG)
			bnew += uint32(convertedB)
			anew += uint32(convertedA)
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
