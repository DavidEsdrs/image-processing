package filters

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/quad"
	"github.com/DavidEsdrs/image-processing/utils"
)

type ContrastFilter struct {
	constrast float64
	logger    *logger.Logger
}

func NewContrastFilter(contrast float64, l *logger.Logger) *ContrastFilter {
	return &ContrastFilter{
		constrast: contrast,
		logger:    l,
	}
}

func (c *ContrastFilter) Execute(img *quad.Quad) error {
	M := c.calculateMFactor(img)

	c.logger.LogProcessf("Contrast: M factor is %v\n", M)

	img.Apply(func(pixel color.RGBA) color.RGBA {
		pixel.R = uint8(utils.Clamp(M+c.constrast*(float64(pixel.R)-M), 0, 255))
		pixel.G = uint8(utils.Clamp(M+c.constrast*(float64(pixel.G)-M), 0, 255))
		pixel.B = uint8(utils.Clamp(M+c.constrast*(float64(pixel.B)-M), 0, 255))
		return pixel
	})

	return nil
}

func (c *ContrastFilter) calculateMFactor(img *quad.Quad) float64 {
	result := 0.0
	pixelCount := img.Cols * img.Rows

	img.Iterate(func(pixel color.RGBA) {
		result += float64(pixel.R)*0.299 + float64(pixel.G)*0.587 + float64(pixel.B)*0.114
	})

	return result / float64(pixelCount)
}
