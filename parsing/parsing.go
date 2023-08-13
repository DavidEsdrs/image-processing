package parsing

import (
	"fmt"
	"image"
	"strings"
)

type ConversionStrategy interface {
	Save(img image.Image, output string) error
}

type ParsingContext struct{}

func NewParsingContext() *ParsingContext {
	return &ParsingContext{}
}

func (cc *ParsingContext) GetConfig(file string) (ConversionStrategy, error) {
	strs := strings.Split(file, ".")
	ftype := strs[len(strs)-1]

	switch ftype {
	case "png":
		return &PngParsingStrategy{}, nil
	case "jpeg":
	case "jpg":
		return &JpgParsingStrategy{}, nil
	}

	return nil, fmt.Errorf("unknown file type")
}
