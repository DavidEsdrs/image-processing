package configs

import (
	"fmt"
	"image"
	"log"
	"strconv"
	"strings"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/models"
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

// config singleton
var config *Config

func (config *Config) ParseConfig(logger logger.Logger, inputImg image.Image) (processor.Processor, error) {
	proc := processor.ImageProcessor{}

	format := strings.Split(config.Output, ".")

	if len(format) <= 1 {
		log.Fatal("Invalid output format")
	}

	config.OutputFormat = format[len(format)-1]

	if config.NearestNeighbor != 1.0 {
		logger.LogProcessf("Resizing image to scale %v - nearest neighbor algorithm\n", config.NearestNeighbor)
		proc.NearestNeighbor(float32(config.NearestNeighbor))
	}
	if config.Grayscale {
		logger.LogProcess("Applying 'grayscale 16 bits' filter")
		proc.Grayscale16()
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
			log.Fatal("wrong arguments count for cropping")
		}

		logger.LogProcessf("Cropping image - arguments: %v, %v, %v, %v", xstart, xend, ystart, yend)
		proc.Crop(xstart, xend, ystart, yend)
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

		model := overlay.ColorModel()

		// convert tensor to the right color model (in that case, the overlay color model is converted to the input color model)
		cc := models.ConverterContext{}

		conv, err := cc.GetConverter(inputImg.ColorModel())

		logger.LogProcessf("Overlay has %v color model", utils.ColorModelString(model))

		if err != nil {
			return nil, err
		}

		tensor := conv.ConvertToModel(overlay)

		config.overlayRect = overlay.Bounds()
		config.backgroundRect = inputImg.Bounds()

		config.ParseOverlayConfigs(&tensor)

		proc.Overlay = &tensor

		proc.SetOverlay(config.DistTop, config.DistRight, config.DistBottom, config.DistLeft)

		logger.LogProcess("Applying overlay")
	}
	if config.Transpose {
		logger.LogProcess("Applying 'transpose' filter")
		proc.Transpose()
	}
	if config.FlipY {
		logger.LogProcess("Applying 'flip Y' filter")
		proc.FlipY()
	}
	if config.FlipX {
		logger.LogProcess("Applying 'flip X' filter")
		proc.FlipX()
	}
	if config.TurnLeft {
		logger.LogProcess("Turning image left - 90 degrees")
		proc.TurnLeft()
	}
	if config.TurnRight {
		logger.LogProcess("Turning image right - 90 degrees")
		proc.TurnRight()
	}

	return &proc, nil
}

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
