package processor

import (
	"fmt"
	"image"
	"image/color"

	"github.com/DavidEsdrs/image-processing/filters"
)

type Overlay *[][]color.Color

type ImageProcessor struct {
	ColorModel     color.Model
	SubsampleRatio image.YCbCrSubsampleRatio
	processes      []Process
	Overlay        Overlay
}

func (ip *ImageProcessor) GetColorModel() color.Model {
	return ip.ColorModel
}

func (ip *ImageProcessor) SetColorModel(colorModel color.Model) {
	ip.ColorModel = colorModel
}

func crop(pImg *[][]color.Color, xstart, xend, ystart, yend int) {
	img := *pImg

	res := make([][]color.Color, yend-ystart)

	for i := range res {
		res[i] = make([]color.Color, xend-xstart)
	}

	for i := ystart; i < yend; i++ {
		for j := xstart; j < xend; j++ {
			res[i-ystart][j-xstart] = img[i][j]
		}
	}

	*pImg = res
}

func (ip *ImageProcessor) Crop(xstart, xend, ystart, yend int) {
	p := func(pImg *[][]color.Color) {
		crop(pImg, xstart, xend, ystart, yend)
	}
	ip.processes = append(ip.processes, p)
}

func flipY(pImg *[][]color.Color) {
	img := *pImg
	rows := len(img)
	cols := len(img[0])

	for i := 0; i < rows/2; i++ {
		for j := 0; j < cols; j++ {
			img[rows-i-1][j], img[i][j] = img[i][j], img[rows-i-1][j]
		}
	}
}

func (ip *ImageProcessor) FlipY() {
	ip.processes = append(ip.processes, flipY)
}

func flipX(pImg *[][]color.Color) {
	img := *pImg
	rows := len(img)
	cols := len(img[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols/2; j++ {
			img[i][cols-j-1], img[i][j] = img[i][j], img[i][cols-j-1]
		}
	}
}

func (ip *ImageProcessor) FlipX() {
	ip.processes = append(ip.processes, flipX)
}

func turnLeft(pImg *[][]color.Color) {
	img := *pImg

	originalRows := len(img)
	originalCols := len(img[0])

	rows := len(img[0])
	cols := len(img)

	res := make([][]color.Color, rows)

	for i := range res {
		res[i] = make([]color.Color, cols)
	}

	for i := 0; i < originalRows; i++ {
		for j := 0; j < originalCols; j++ {
			res[rows-j-1][i] = img[i][j]
		}
	}

	*pImg = res
}

func (ip *ImageProcessor) TurnLeft() {
	ip.processes = append(ip.processes, turnLeft)
}

func turnRight(pImg *[][]color.Color) {
	img := *pImg

	originalRows := len(img)
	originalCols := len(img[0])

	rows := len(img[0])
	cols := len(img)

	res := make([][]color.Color, rows)

	for i := range res {
		res[i] = make([]color.Color, cols)
	}

	for i := 0; i < originalRows; i++ {
		for j := 0; j < originalCols; j++ {
			res[j][originalRows-i-1] = img[i][j]
		}
	}

	*pImg = res
}

func (ip *ImageProcessor) TurnRight() {
	ip.processes = append(ip.processes, turnRight)
}

func transpose(pImg *[][]color.Color) {
	img := *pImg

	res := make([][]color.Color, len(img[0]))

	for i := range res {
		res[i] = make([]color.Color, len(img))
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			res[j][i] = img[i][j]
		}
	}

	*pImg = res
}

func (ip *ImageProcessor) Transpose() {
	ip.processes = append(ip.processes, transpose)
}

func nearestNeighbor(pImg *[][]color.Color, factor float32) {
	img := *pImg

	if factor <= 0 {
		fmt.Printf("factor is less than or equal to 0 - resize won't happen")
		return
	}

	proportion := int(1 / factor)
	rows := int(len(img) / proportion)
	cols := int(len(img[0]) / proportion)

	res := make([][]color.Color, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]color.Color, cols)

		for j := 0; j < cols; j++ {
			res[i][j] = img[i*proportion][j*proportion]
		}
	}

	*pImg = res
}

// Basic resize operation. Applies nearest neighbor algorithm
func (ip *ImageProcessor) NearestNeighbor(factor float32) {
	p := func(pImg *[][]color.Color) {
		nearestNeighbor(pImg, factor)
	}
	ip.processes = append(ip.processes, p)
}

func getAdjacentPixels(pImg *[][]color.Color, x int, y int) [4]color.Color {
	img := *pImg
	var res [4]color.Color
	// get pixels around target pixel
	res[0] = img[x-1][y]
	res[1] = img[x][y+1]
	res[2] = img[x+1][y]
	res[3] = img[x][y-1]
	return res
}

func (ip *ImageProcessor) Grayscale16() {
	bAw := func(pImg *[][]color.Color) {
		filters.Grayscale16(pImg, ip)
	}
	ip.processes = append(ip.processes, bAw)
}

func overlay(pImg *[][]color.Color, pOverlay *[][]color.Color, distTop, distRight, distBottom, distLeft int) {
	img := *pImg
	overlay := *pOverlay

	rows := len(overlay)
	cols := len(overlay[0])

	imgRows := len(img)
	imgCols := len(img[0])

	for y := 0; y < rows && y+distTop < imgRows; y++ {
		for x := 0; x < cols && x+distLeft < imgCols; x++ {
			img[y+distTop][x+distLeft] = overlay[y][x]
		}
	}

	*pImg = img
}

func (ip *ImageProcessor) SetOverlay(distTop, distRight, distBottom, distLeft int) {
	f := func(pImg *[][]color.Color) {
		overlay(pImg, ip.Overlay, distTop, distRight, distBottom, distLeft)
	}
	ip.processes = append(ip.processes, f)
}

func (ip *ImageProcessor) Execute(source *[][]color.Color) [][]color.Color {
	for _, process := range ip.processes {
		process(source)
	}
	return *source
}
