package utils

import (
	"image"
	"image/color"
	"math"
	"os"

	"github.com/DavidEsdrs/image-processing/quad"
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
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func Min[T Number](a, b, c T) T {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

// clamp ensures the value stays within the specified range.
func Clamp(value float64, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
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

func ConvertIntoQuad(img image.Image) *quad.Quad {
	size := img.Bounds().Size()
	pixels := quad.NewQuad(size.X, size.Y)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r >>= 8
			g >>= 8
			b >>= 8
			a >>= 8
			pixels.SetPixel(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	return pixels
}

// instead of recreating the whole []uint8, we try to get it from the image first
func GetQuad(img image.Image) *quad.Quad {
	size := img.Bounds().Size()
	pixels := quad.NewEmptyQuad(size.X, size.Y)

	switch img := img.(type) {
	case *image.NRGBA:
		pixels.SetSlice(img.Pix)
	case *image.RGBA:
		pixels.SetSlice(img.Pix)
	case *image.Gray, *image.Alpha, *image.Alpha16, *image.NRGBA64, *image.RGBA64,
		*image.Gray16, *image.CMYK, *image.Paletted, *image.YCbCr:
		return ConvertIntoQuad(img)
	}

	return pixels
}

// Kernel represents a gaussian kernel in a linear structure
type Kernel struct {
	Size int
	Data []float64
}

func (k *Kernel) GetValue(x, y int) float64 {
	index := y*k.Size + x
	return k.Data[index]
}

// Creates a gaussian kernel of the given size, using the given sigma
func GaussianKernel(size int, sigma float64) Kernel {
	if size%2 == 0 {
		size++
	}

	center := size / 2
	kernelData := make([]float64, size*size)
	sum := 0.0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			x := float64(j - center)
			y := float64(i - center)
			index := i*size + j
			kernelData[index] = math.Exp(-(x*x+y*y)/(2*sigma*sigma)) / (2 * math.Pi * sigma * sigma)
			sum += kernelData[index]
		}
	}

	for i := 0; i < size*size; i++ {
		kernelData[i] /= sum
	}

	return Kernel{Size: size, Data: kernelData}
}
