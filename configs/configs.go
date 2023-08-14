package configs

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/DavidEsdrs/image-processing/processor"
)

type Config struct {
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

	Ssr            int
	SubsampleRatio image.YCbCrSubsampleRatio
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
		proc.BlackAndWhite()
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

	return &proc
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
