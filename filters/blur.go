package filters

import (
	"image/color"
	"sync"
	"time"

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
	start := time.Now()

	height := len(*tensor)
	width := len((*tensor)[0])

	copy := deepCopy(tensor)

	var wg sync.WaitGroup

	process := func(x, y int) {
		defer wg.Done()
		_, _, _, a := (*tensor)[y][x].RGBA()
		r, g, b := bf.getValuesForPixel(tensor, &copy, x, y)
		(*tensor)[y][x] = color.RGBA{
			R: r,
			G: g,
			B: b,
			A: uint8(a >> 8),
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			wg.Add(1)
			go process(x, y)
		}
	}

	wg.Wait()

	duration := time.Since(start)
	bf.l.LogProcessf("it took %vms to process the blur", duration.Milliseconds())

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
) (r, g, b uint8) {
	height := len(*tensor)
	width := len((*tensor)[0])

	var (
		rnew uint8
		gnew uint8
		bnew uint8
	)

	sy := fixed(startY - (bf.kernelSize / 2))
	sx := fixed(startX - (bf.kernelSize / 2))

	endY := fixed(startY + (bf.kernelSize / 2))
	endX := fixed(startX + (bf.kernelSize / 2))

	for y := sy; y <= endY && y < height; y++ {
		for x := sx; x <= endX && x < width; x++ {
			r, g, b, _ := (*copy)[y][x].RGBA()

			rc := uint8(r >> 8)
			gc := uint8(g >> 8)
			bc := uint8(b >> 8)

			w := bf.kernel[y-sy][x-sx]

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

func fixed(x int) int {
	if x < 0 {
		return 0
	}
	return x
}
