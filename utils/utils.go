package utils

import (
	"image"
	"image/color"
	"math"
	"os"
)

type Number interface {
	int | float32 | float64 | uint32 | uint64
}

func pow[T Number](base T, pow int) T {
	var res T = 1
	for k := 0; k < pow; k++ {
		res *= base
	}
	return res
}

func cubic(x float64) float64 {
	absX := math.Abs(x)
	if absX <= 1 {
		return 1 - 2*pow(absX, 2) + pow(absX, 3)
	} else if absX > 1 && absX <= 2 {
		return 4 - 8*absX + 5*pow(absX, 2) - pow(absX, 3)
	} else {
		return 0
	}
}

func normalize(kernel [][]float64, size int) [][]float64 {
	sum := 0.0
	// get the sum of the elements in the kernel
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += kernel[i][j]
		}
	}
	// divide each element by the total sum
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			kernel[i][j] /= sum
		}
	}
	return kernel
}

func CreateBicubicKernel(size int) [][]float64 {
	kernel := make([][]float64, size)
	for i := range kernel {
		kernel[i] = make([]float64, size)
	}

	// Fill kernel with cubic interpolation
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			x := (float64(j) - float64(size)/2 + 0.5) / (float64(size) / 2)
			y := (float64(i) - float64(size)/2 + 0.5) / (float64(size) / 2)
			kernel[i][j] = cubic(x) * cubic(y)
		}
	}

	// Normalize
	kernel = normalize(kernel, size)

	return kernel
}

func Max[T Number](a, b, c T) T {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func Min[T Number](a, b, c T) T {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// Load image file with the given path - return an error and image.Image nil
// if it fails
func LoadImage(file string) (image.Image, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// Convert the image into a tensor to further manipulation
func ConvertIntoTensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	pixels := make([][]color.Color, size.Y)

	for y := 0; y < size.Y; y++ {
		pixels[y] = make([]color.Color, size.X)
		for x := 0; x < size.X; x++ {
			pixels[y][x] = img.At(x, y)
		}
	}

	return pixels
}

// Creates a gaussian kernel of the given size, using the given sigma
func GaussianKernel(size int, sigma float64) [][]float64 {
	if size%2 == 0 {
		size++
	}

	kernel := make([][]float64, size)
	center := size / 2
	sum := 0.0

	for i := 0; i < size; i++ {
		kernel[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			x := float64(j - center)
			y := float64(i - center)
			kernel[i][j] = math.Exp(-(x*x+y*y)/(2*sigma*sigma)) / (2 * math.Pi * sigma * sigma)
			sum += kernel[i][j]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			kernel[i][j] /= sum
		}
	}

	return kernel
}
