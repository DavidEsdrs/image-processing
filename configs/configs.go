package configs

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/DavidEsdrs/image-processing/processor"
)

type Config struct {
	// Filters
	Input           string
	Output          string
	FlipY           bool
	FlipX           bool
	Transpose       bool
	Grayscale       bool
	TurnLeft        bool
	TurnRight       bool
	NearestNeighbor float64
	Crop            string
	Overlay         string

	// YCbCr
	Ssr            int
	SubsampleRatio image.YCbCrSubsampleRatio

	// JPEG
	Quality int

	OutputFormat string

	// Overlay
	DistTop    int
	DistLeft   int
	DistBottom int
	DistRight  int
	Fill       bool
}

func (cfg *Config) SetSubsampleRatio(ratio int) {
	switch ratio {
	case 444, 5:
		cfg.SubsampleRatio = image.YCbCrSubsampleRatio444
		return
	case 422, 4:
		cfg.SubsampleRatio = image.YCbCrSubsampleRatio422
		return
	case 420, 3:
		cfg.SubsampleRatio = image.YCbCrSubsampleRatio420
		return
	case 440, 2:
		cfg.SubsampleRatio = image.YCbCrSubsampleRatio440
		return
	case 411, 1:
		cfg.SubsampleRatio = image.YCbCrSubsampleRatio411
		return
	case 410, 0:
		cfg.SubsampleRatio = image.YCbCrSubsampleRatio410
		return
	}
	cfg.SubsampleRatio = image.YCbCrSubsampleRatio444
	fmt.Printf("Invalid subsample ratio: %v - default to 4:4:4\n", ratio)
}

var config *Config

func (config *Config) ParseConfig() processor.Processor {
	proc := processor.ImageProcessor{}

	format := strings.Split(config.Output, ".")

	if len(format) <= 1 {
		log.Fatal("Invalid output format")
	}

	config.OutputFormat = format[len(format)-1]

	if config.Transpose {
		proc.Transpose()
	}
	if config.FlipY {
		proc.FlipY()
	}
	if config.FlipX {
		proc.FlipX()
	}
	if config.NearestNeighbor != 1.0 {
		proc.NearestNeighbor(float32(config.NearestNeighbor))
	}
	if config.Grayscale {
		proc.Grayscale16()
	}
	if config.TurnLeft {
		proc.TurnLeft()
	}
	if config.TurnRight {
		proc.TurnRight()
	}
	if config.Crop != "" {
		str := strings.Split(config.Crop, ",")

		var xstart int
		var xend int
		var ystart int
		var yend int

		if len(str) == 4 {
			xstart, _ = strconv.Atoi(str[0])
			xend, _ = strconv.Atoi(str[1])
			ystart, _ = strconv.Atoi(str[2])
			yend, _ = strconv.Atoi(str[3])
		} else {
			xend, _ = strconv.Atoi(str[0])
			yend, _ = strconv.Atoi(str[1])
		}

		proc.Crop(xstart, xend, ystart, yend)
	}
	if config.Ssr != 0 {
		config.SetSubsampleRatio(config.Ssr)
	}
	if config.Quality > 100 || config.Quality < 0 {
		fmt.Printf("quality value too high or too low - default to 100\n")
		config.Quality = 100
	}
	if config.Overlay != "" {
		imgFile, err := os.Open(config.Overlay)
		if err != nil {
			log.Fatal("Can't open overlay file")
		}
		img, _, err := image.Decode(imgFile)
		if err != nil {
			log.Fatal("Can't decode overlay file")
		}
		tensor := ConvertIntoTensor(img)
		proc.Overlay = &tensor
		proc.SetOverlay(config.DistTop, config.DistRight, config.DistBottom, config.DistLeft)
		imgFile.Close()
	}

	return &proc
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

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
