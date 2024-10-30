package filters

import (
	"image/color"
	"sync"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/quad"
	"github.com/DavidEsdrs/image-processing/utils"
)

type BlurFilter struct {
	l          logger.Logger
	sigma      float64
	kernelSize int // this will define how blurry the image is
	kernel     utils.Kernel
}

func NewBlurFilter(l logger.Logger, sigma float64, kernelSize int) (BlurFilter, error) {
	kernel := utils.GaussianKernel(kernelSize, sigma)
	if sigma == 1 {
		sigma = float64(kernelSize) / 2
	}
	bf := BlurFilter{l: l, sigma: sigma, kernelSize: kernelSize, kernel: kernel}
	return bf, nil
}

func (bf BlurFilter) Execute(img *quad.Quad) error {
	height := img.Rows
	width := img.Cols

	paddingSize := bf.kernelSize / 2
	copy := padImage(img, paddingSize)

	var wg sync.WaitGroup

	process := func(x, y int) {
		defer wg.Done()
		r, g, b, a := bf.getValuesForPixel(img, copy, x, y)

		img.SetPixel(x, y, color.RGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: uint8(a >> 8),
		})
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

func padImage(img *quad.Quad, paddingSize int) *quad.Quad {
	width := img.Cols
	targetWidth := img.Cols + paddingSize*2

	height := img.Rows
	targetHeight := img.Rows + paddingSize*2

	output := quad.NewQuad(targetWidth, targetHeight)

	padHorizontal := func() {
		for y := 0; y < height; y++ {
			firstPixel := img.GetPixel(0, y)
			lastPixel := img.GetPixel(width-1, y)

			for x := 0; x < targetWidth; x++ {
				if x > paddingSize && x < targetWidth-paddingSize {
					output.SetPixel(x, y, img.GetPixel(x-paddingSize, y))
				} else if x < paddingSize {
					output.SetPixel(x, y, firstPixel)
				} else {
					output.SetPixel(x, y, lastPixel)
				}
			}
		}
	}

	padHorizontal()

	padVertical := func() *quad.Quad {
		newOutput := quad.NewQuad(targetWidth, targetHeight)

		firstLine := output.GetRow(0)
		lastLine := output.GetRow(height - 1)

		for y := 0; y < targetHeight; y++ {
			if y > paddingSize && y < targetHeight-paddingSize {
				newOutput.SetRow(y, output.GetRow(y-paddingSize))
			} else if y < paddingSize {
				newOutput.SetRow(y, firstLine)
			} else {
				newOutput.SetRow(y, lastLine)
			}
		}

		return newOutput
	}

	return padVertical()
}

func (bf BlurFilter) getValuesForPixel(
	tensor *quad.Quad,
	copy *quad.Quad,
	startX,
	startY int,
) (r, g, b, a uint32) {
	height := tensor.Rows
	width := tensor.Cols

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
			w := bf.kernel.GetValue(x-sx, y-sy)
			r, g, b, a := copy.GetPixel(x, y).RGBA()

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
