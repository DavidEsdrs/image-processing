package filters

import (
	"image/color"
	"sync"

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

	paddingSize := bf.kernelSize / 2
	copy := padImage(tensor, paddingSize)

	var wg sync.WaitGroup

	process := func(x, y int) {
		defer wg.Done()
		r, g, b, a := bf.getValuesForPixel(tensor, copy, x, y)
		(*tensor)[y][x] = color.RGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: uint8(a >> 8),
		}
	}

	for y := paddingSize; y < height+paddingSize; y++ {
		for x := paddingSize; x < width+paddingSize; x++ {
			wg.Add(1)
			go process(x-paddingSize, y-paddingSize)
		}
	}

	wg.Wait()

	return nil
}

func padImage(img *[][]color.Color, paddingSize int) *[][]color.Color {
	width := len((*img)[0])
	targetWidth := len((*img)[0]) + paddingSize*2

	height := len(*img)
	targetHeight := len(*img) + paddingSize*2

	output := make([][]color.Color, height)

	for y := 0; y < height; y++ {
		output[y] = make([]color.Color, targetWidth)
	}

	padHorizontal := func() {
		for y := 0; y < height; y++ {
			firstPixel := (*img)[y][0]
			lastPixel := (*img)[y][width-1]

			for x := 0; x < targetWidth; x++ {
				if x > paddingSize && x < targetWidth-paddingSize {
					output[y][x] = (*img)[y][x-paddingSize]
				} else if x < paddingSize {
					output[y][x] = firstPixel
				} else {
					output[y][x] = lastPixel
				}
			}
		}
	}

	padHorizontal()

	padVertical := func() *[][]color.Color {
		newOutput := make([][]color.Color, targetHeight)

		firstLine := output[0]
		lastLine := output[height-1]

		for y := 0; y < targetHeight; y++ {
			if y > paddingSize && y < targetHeight-paddingSize {
				newOutput[y] = output[y-paddingSize]
			} else if y < paddingSize {
				newOutput[y] = firstLine
			} else {
				newOutput[y] = lastLine
			}
		}

		return &newOutput
	}

	return padVertical()
}

func (bf BlurFilter) getValuesForPixel(
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
			w := bf.kernel[y-sy][x-sx]
			r, g, b, a := (*copy)[y][x].RGBA()

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
