// This package is used to transform a image from a color model to another.
//
// The main difference of this package with the "convert" package is that the
// strategies of this package converts from a image.Image to a [][]color.Color,
// while the strategies from "convert" converts from [][]color.Color to image.Image
//
// You might wonder why these strategies weren't implemented in "convert" package.
// This is to avoid circular imports, as the "configs" package would have to
// import "convert" package, which already imports "configs" package, so, it would
// be impossible.
//

package models

import (
	"fmt"
	"image/color"

	"github.com/DavidEsdrs/image-processing/interfaces"
)

type ConverterContext struct{}

func (cc ConverterContext) GetConverter(model color.Model) (interfaces.Converter, error) {
	switch model {
	case color.Alpha16Model:
		return &Alpha16Strategy{}, nil
	case color.AlphaModel:
		return &AlphaStrategy{}, nil
	case color.CMYKModel:
		return &CmykStrategy{}, nil
	case color.Gray16Model:
		return &Gray16Strategy{}, nil
	case color.GrayModel:
		return &GrayStrategy{}, nil
	case color.NRGBA64Model:
		return &Nrgba64Strategy{}, nil
	case color.NRGBAModel:
		return &NrgbaStrategy{}, nil
	case color.RGBA64Model:
		return &Rgba64Strategy{}, nil
	case color.RGBAModel:
		return &RgbaStrategy{}, nil
	case color.YCbCrModel:
		return &YcbcrStrategy{}, nil
	case color.NYCbCrAModel:
		return nil, fmt.Errorf("unsupported color model - No support for NYCbCr yet")
	}

	return nil, fmt.Errorf("unsupported color model - No support for Paletted images as background")
}
