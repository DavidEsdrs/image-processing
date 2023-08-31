package configs

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/DavidEsdrs/image-processing/filters"
	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/processor"
	"github.com/DavidEsdrs/image-processing/utils"
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
	DistTop        int
	DistLeft       int
	DistBottom     int
	DistRight      int
	Fill           bool
	overlayRect    image.Rectangle
	backgroundRect image.Rectangle
}

// config singleton
var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			DistTop:    0,
			DistLeft:   0,
			DistBottom: -1,
			DistRight:  -1,
			Fill:       false,
		}
	}
	return config
}

func (config *Config) ParseConfig(logger logger.Logger, inputImg image.Image) (*processor.Invoker, error) {
	invoker := processor.Invoker{}

	format := strings.Split(config.Output, ".")

	if len(format) <= 1 {
		return nil, fmt.Errorf("invalid output format")
	}

	config.OutputFormat = format[len(format)-1]

	if config.NearestNeighbor != 1.0 {
		if config.NearestNeighbor <= 0 {
			return nil, fmt.Errorf("invalid scale factor to nearest neighbor")
		}
		logger.LogProcessf("Resizing image to scale %v - nearest neighbor algorithm\n", config.NearestNeighbor)

		f := filters.NewNearestNeighborFilter(float32(config.NearestNeighbor))

		invoker.AddProcess(f)
	}
	if config.Grayscale {
		logger.LogProcess("Applying 'grayscale 16 bits' filter")
		f := filters.NewGrayscale16Filter(&invoker)
		invoker.AddProcess(f)
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
		} else if len(str) == 2 {
			xend, _ = strconv.Atoi(str[0])
			yend, _ = strconv.Atoi(str[1])
		} else {
			return nil, fmt.Errorf("wrong arguments count for cropping")
		}

		f, err := filters.NewCropFilter(inputImg, xstart, xend, ystart, yend)

		if err != nil {
			return nil, err
		}

		logger.LogProcessf("Cropping image - arguments: %v, %v, %v, %v", xstart, xend, ystart, yend)
		invoker.AddProcess(f)
	}
	if config.Ssr != 0 {
		logger.LogProcessf("Changing subsampling ratio - using %v\n", config.Ssr)
		config.SetSubsampleRatio(config.Ssr)
	}
	if config.Quality > 100 || config.Quality < 0 {
		logger.LogWarn("quality value too high or too low - default to 100\n")
		config.Quality = 100
	}
	if config.Overlay != "" {
		overlay, err := utils.LoadImage(config.Overlay)

		if err != nil {
			return nil, err
		}

		f, err := filters.NewOverlayFilter(logger, overlay, inputImg, config.DistTop, config.DistRight, config.DistLeft, config.DistBottom)

		if err != nil {
			return nil, err
		}

		logger.LogProcess("Applying overlay")

		invoker.AddProcess(f)
	}
	if config.Transpose {
		logger.LogProcess("Applying 'transpose' filter")
		f := filters.NewTransposeFilter()
		invoker.AddProcess(f)
	}
	if config.FlipY {
		logger.LogProcess("Applying 'flip Y' filter")
		f := filters.NewFlipYFilter()
		invoker.AddProcess(f)
	}
	if config.FlipX {
		logger.LogProcess("Applying 'flip X' filter")
		f := filters.NewFlipXFilter()
		invoker.AddProcess(f)
	}
	if config.TurnLeft {
		logger.LogProcess("Turning image left - 90 degrees")
		f := filters.NewTurnLeftFilter()
		invoker.AddProcess(f)
	}
	if config.TurnRight {
		logger.LogProcess("Turning image right - 90 degrees")
		f := filters.NewTurnRightFilter()
		invoker.AddProcess(f)
	}

	return &invoker, nil
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
