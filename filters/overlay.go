package filters

import (
	"image"
	"image/color"
	"math"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/models"
	"github.com/DavidEsdrs/image-processing/utils"
)

type OverlayFilter struct {
	overlay                                  *[][]color.Color
	distTop, distRight, distLeft, distBottom int
	overlayRect                              image.Rectangle
	backgroundRect                           image.Rectangle
}

func NewOverlayFilter(logger logger.Logger, overlay, background image.Image, distTop, distRight, distLeft, distBottom int) (OverlayFilter, error) {
	cc := models.ConverterContext{}
	conv, err := cc.GetConverter(background.ColorModel())

	logger.LogProcessf("Overlay has %v color model", utils.ColorModelString(overlay.ColorModel()))

	if err != nil {
		return OverlayFilter{}, err
	}

	overlayRect := overlay.Bounds()
	backgroundRect := background.Bounds()

	tensor := conv.ConvertToModel(overlay)

	ovf := OverlayFilter{
		overlay:        &tensor,
		distTop:        distTop,
		distRight:      distRight,
		distLeft:       distLeft,
		distBottom:     distBottom,
		overlayRect:    overlayRect,
		backgroundRect: backgroundRect,
	}

	ovf.parseOverlayConfigs(&tensor)

	return ovf, nil
}

func (ovf OverlayFilter) Execute(tensor *[][]color.Color) error {
	img := *tensor
	overlay := *ovf.overlay

	rows := len(overlay)
	cols := len(overlay[0])

	imgRows := len(img)
	imgCols := len(img[0])

	for y := 0; y < rows && y+ovf.distTop < imgRows; y++ {
		for x := 0; x < cols && x+ovf.distLeft < imgCols; x++ {
			img[y+ovf.distTop][x+ovf.distLeft] = overlay[y][x]
		}
	}

	*tensor = img
	return nil
}

func (ovf *OverlayFilter) parseOverlayConfigs(tensor *[][]color.Color) {
	if ovf.distRight != math.MinInt32 {
		ovf.parseHorizontalAxis()
	}
	if ovf.distBottom != math.MinInt32 {
		ovf.parseVerticalAxis()
	}
	if ovf.distTop < 0 {
		ovf.adjustOverlayVerticalOffset(tensor)
	}
	if ovf.distLeft < 0 {
		ovf.adjustOverlayHorizontalOffset(tensor)
	}
}

func (cfg *OverlayFilter) parseHorizontalAxis() {
	ovPlustDist := cfg.distRight + cfg.overlayRect.Max.X
	distToLeft := cfg.backgroundRect.Max.X - ovPlustDist
	cfg.distLeft = distToLeft
}

func (cfg *OverlayFilter) parseVerticalAxis() {
	ovPlustDist := cfg.distBottom + cfg.overlayRect.Max.Y
	distToTop := cfg.backgroundRect.Max.Y - ovPlustDist
	cfg.distTop = distToTop
}

func (ovf *OverlayFilter) adjustOverlayVerticalOffset(tensor *[][]color.Color) {
	absOffset := int(math.Abs(float64(ovf.distTop)))
	pTensor := (*tensor)[absOffset:]
	*tensor = pTensor
	ovf.distTop = 0
}

func (ovf *OverlayFilter) adjustOverlayHorizontalOffset(tensor *[][]color.Color) {
	absOffset := int(math.Abs(float64(ovf.distLeft)))
	pTensor := slice2D(*tensor, 0, absOffset)
	*tensor = pTensor
	ovf.distLeft = 0
}

func slice2D(slice [][]color.Color, startRow int, startCol int) [][]color.Color {
	result := make([][]color.Color, len(slice)-startRow)
	for i := startRow; i < len(slice); i++ {
		result[i-startRow] = slice[i][startCol:]
	}
	return result
}
