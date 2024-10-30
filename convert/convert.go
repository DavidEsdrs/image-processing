// This package determines the input image's color model and provides conversion
// strategies for transforming tensors ([][]color.Color) back into images using
// their respective color models.
//
// Each strategy implements a method "Convert" to convert the tensor ([][]color.Color)
// back to an image as its own color model. For instance, the NRGBA64 strategy
// will read each color.Color from the tensor as a NRGBA64 pixel and returning it
// as an image.Image.
//
// This approach, ensures that the images will be converted back to the original
// color model, i.e the original color model will be respected

package convert

import (
	"fmt"
	"image"
	"image/color"

	"github.com/DavidEsdrs/image-processing/configs"
	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/palette"
	"github.com/DavidEsdrs/image-processing/quad"
)

type ConversionStrategy interface {
	Convert(pixels *quad.Quad) image.Image
}

type ConversionContext struct {
	logger *logger.Logger
}

func NewConversionContext(logger *logger.Logger) *ConversionContext {
	return &ConversionContext{logger}
}

func (cc *ConversionContext) GetConversor(img image.Image, mdl color.Model) (ConversionStrategy, error) {
	var model color.Model
	if mdl == nil {
		model = img.ColorModel()
	} else {
		model = mdl
	}
	switch model {
	case color.Alpha16Model:
		cc.logger.LogProcess("Given image has Alpha 16 bits color Model")
		return &Alpha16Strategy{}, nil
	case color.AlphaModel:
		cc.logger.LogProcess("Given image has Alpha 8 bits color Model")
		return &AlphaStrategy{}, nil
	case color.CMYKModel:
		cc.logger.LogProcess("Given image has CMYK color Model")
		return &CmykStrategy{}, nil
	case color.Gray16Model:
		cc.logger.LogProcess("Given image has Gray 16 bits color Model")
		return &Gray16Strategy{}, nil
	case color.GrayModel:
		cc.logger.LogProcess("Given image has Gray 8 bits color Model")
		return &GrayStrategy{}, nil
	case color.NRGBA64Model:
		cc.logger.LogProcess("Given image has NRGBA 64 bits color Model")
		return &Nrgba64Strategy{}, nil
	case color.NRGBAModel:
		cc.logger.LogProcess("Given image has NRGBA 32 bits color Model")
		return &NrgbaStrategy{}, nil
	case color.RGBA64Model:
		cc.logger.LogProcess("Given image has RGBA 64 bits color Model")
		return &Rgba64Strategy{}, nil
	case color.RGBAModel:
		cc.logger.LogProcess("Given image has RGBA 32 bits color Model")
		return &RgbaStrategy{}, nil
	case color.YCbCrModel:
		cc.logger.LogProcess("Given image has YCbCr color Model")
		cfg := configs.GetConfig()

		// if the user passed a custom subsampling ratio, use it
		if cfg.Ssr != 0 {
			return &YcbcrStrategy{cfg.SubsampleRatio}, nil
		}

		if imgYcbcr, ok := img.(*image.YCbCr); ok {
			subsamplingRatio := imgYcbcr.SubsampleRatio
			return &YcbcrStrategy{subsamplingRatio}, nil
		}

		// Assert
		return nil, fmt.Errorf("unsupported color model")
	case color.NYCbCrAModel:
		cc.logger.LogProcess("Given image has NYCbCr color Model")
		return nil, fmt.Errorf("unsupported color model")
	}

	cc.logger.LogProcess("Given image has a custom palette color Model")

	// TODO: Add flag to ignore unknown color models
	noIgnoreUnknown := true
	if noIgnoreUnknown {
		plt, _ := palette.GetPalette(img)
		return &PaletteStrategy{plt}, nil
	}

	return nil, fmt.Errorf("unsupported color model")
}
