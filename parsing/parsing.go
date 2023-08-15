package parsing

import (
	"fmt"
	"image"
)

type ConversionStrategy interface {
	Save(img image.Image, output string) error
}

type ParsingContext struct{}

func NewParsingContext() *ParsingContext {
	return &ParsingContext{}
}

func (cc *ParsingContext) GetConfig(format string) (ConversionStrategy, error) {
	switch format {
	case "png":
		return &PngParsingStrategy{}, nil
	case "jpeg", "jpg":
		return &JpgParsingStrategy{}, nil
	}

	return nil, fmt.Errorf("unknown file type")
}
