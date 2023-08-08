package utils

import "math"

type Number interface {
	int | float32 | float64
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
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
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

func Max(a, b, c uint32) uint32 {
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

func Min(a, b, c uint32) uint32 {
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
