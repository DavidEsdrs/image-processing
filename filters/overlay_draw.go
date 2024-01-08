package filters

import (
	"image"
	"image/color"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/utils"
)

type OverlayDrawFilter struct {
	logger                                   logger.Logger
	overlay                                  image.Image
	background                               image.Image
	distTop, distRight, distLeft, distBottom int
	overlayRect                              image.Rectangle
	backgroundRect                           image.Rectangle
	fillBackground                           bool
}

func NewOverlayDrawFilter(
	logger logger.Logger,
	overlay image.Image,
	background image.Image,
	distTop, distRight, distLeft, distBottom int,
	overlayRect image.Rectangle,
	backgroundRect image.Rectangle,
	fillBackground bool,
) OverlayDrawFilter {
	ovd := OverlayDrawFilter{
		logger,
		overlay,
		background,
		distTop,
		distRight,
		distLeft,
		distBottom,
		overlayRect,
		backgroundRect,
		fillBackground,
	}
	return ovd
}

func __NewOverlayDrawFilter(
	logger logger.Logger,
	overlay image.Image,
	background image.Image,
	distTop, distRight, distLeft, distBottom int,
	fillBackground bool,
) OverlayDrawFilter {
	ovd := OverlayDrawFilter{
		logger:         logger,
		overlay:        overlay,
		background:     background,
		distTop:        distTop,
		distRight:      distRight,
		distLeft:       distLeft,
		distBottom:     distBottom,
		fillBackground: fillBackground,
	}
	return ovd
}

func getBackgroundSizes(bg, ov image.Image, dtop, dbottom, dleft, dright int) image.Rectangle {
	bgSize := bg.Bounds()
	ovSize := ov.Bounds()
	result := image.Rectangle{}

	result.Max.Y = utils.Max(bgSize.Max.Y, dtop+ovSize.Max.Y, dbottom+ovSize.Max.Y)
	result.Max.X = utils.Max(bgSize.Max.X, dleft+ovSize.Max.X, dright+ovSize.Max.X)

	return result
}

func (ovd OverlayDrawFilter) Execute(tensor *[][]color.Color) error {
	// newTensor := createBackgroundTensorFromSize(ovd.backgroundRect)

	return nil
}

func createBackgroundTensorFromSize(backgroundRect image.Rectangle) *[][]color.Color {
	width := backgroundRect.Max.X
	height := backgroundRect.Max.Y

	newTensor := make([][]color.Color, height)

	for y := 0; y < width; y++ {
		newTensor[y] = make([]color.Color, width)
	}

	return &newTensor
}
