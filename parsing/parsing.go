package parsing

import (
	"fmt"
	"image"

	"github.com/DavidEsdrs/image-processing/configs"
	"github.com/DavidEsdrs/image-processing/logger"
)

type ConversionStrategy interface {
	Save(img image.Image, output string) error
}

type ParsingContext struct {
	logger *logger.Logger
}

func NewParsingContext(logger *logger.Logger) *ParsingContext {
	return &ParsingContext{logger}
}

func (cc *ParsingContext) GetConfig() (ConversionStrategy, error) {
	format := configs.GetConfig().OutputFormat

	switch format {
	case "png":
		cc.logger.LogProcess("Converting/parsing as PNG")
		return &PngParsingStrategy{logger: cc.logger}, nil
	case "jpeg", "jpg":
		cc.logger.LogProcess("Converting/parsing as JPEG")
		return &JpgParsingStrategy{logger: cc.logger}, nil
	}

	return nil, fmt.Errorf("unknown file type")
}
