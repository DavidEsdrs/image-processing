package processor

import (
	"fmt"
	"image/color"

	"github.com/DavidEsdrs/image-processing/utils"
)

type Processor interface {
	Crop(xstart, xend, ystart, yend int)
	FlipX()
	FlipY()
	TurnLeft()
	TurnRight()
	Transpose()
	BlackAndWhite()
	NearestNeighbor(factor float32)
	Execute(source *[][]color.Color) [][]color.Color
}

type Process func(*[][]color.Color)

type ImageProcessor struct {
	processes []Process
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

func (ip *ImageProcessor) BlackAndWhite() {
	ip.processes = append(ip.processes, blackAndWhite)
}

func blackAndWhite(pImg *[][]color.Color) {
	img := *pImg

	var rows int
	var cols int

	rows = int(len(img))
	cols = int(len(img[0]))

	res := make([][]color.Color, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]color.Color, cols)

		for j := 0; j < cols; j++ {
			originalColor := img[i][j]

			r, g, b, _ := originalColor.RGBA()

			maxValue := uint32(utils.Max(r, g, b))
			minValue := uint32(utils.Min(r, g, b))

			grayValue := uint8((maxValue + minValue) / 2 >> 8)

			newColor := color.RGBA{grayValue, grayValue, grayValue, 255}

			res[i][j] = newColor
		}
	}

	*pImg = res
}

func (ip *ImageProcessor) Execute(source *[][]color.Color) [][]color.Color {
	for _, process := range ip.processes {
		process(source)
	}
	return *source
}
