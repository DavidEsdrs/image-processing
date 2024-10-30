package filters

import (
	"image"
	"math"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/quad"
	"github.com/DavidEsdrs/image-processing/utils"
)

type OverlayFilter struct {
	logger                                   logger.Logger
	overlay                                  *quad.Quad
	distTop, distRight, distLeft, distBottom int
	overlayRect                              image.Rectangle
	backgroundRect                           image.Rectangle
}

func NewOverlayFilter(logger logger.Logger, overlay, background image.Image, distTop, distRight, distLeft, distBottom int) (OverlayFilter, error) {
	overlayRect := overlay.Bounds()
	backgroundRect := background.Bounds()

	q := utils.ConvertIntoQuad(overlay)

	ovf := OverlayFilter{
		overlay:        q,
		distTop:        distTop,
		distRight:      distRight,
		distLeft:       distLeft,
		distBottom:     distBottom,
		overlayRect:    overlayRect,
		backgroundRect: backgroundRect,
		logger:         logger,
	}

	ovf.parseOverlayConfigs(q)

	return ovf, nil
}

func (ovf OverlayFilter) Execute(img *quad.Quad) error {
	overlay := ovf.overlay

	rows := overlay.Rows
	cols := overlay.Cols

	imgRows := img.Rows
	imgCols := img.Cols

	for y := 0; y < rows && y+ovf.distTop < imgRows; y++ {
		for x := 0; x < cols && x+ovf.distLeft < imgCols; x++ {
			overlayPixel := overlay.GetPixel(x, y)
			img.SetPixel(x+ovf.distLeft, y+ovf.distTop, overlayPixel)
		}
	}

	return nil
}

func (ovf *OverlayFilter) parseOverlayConfigs(q *quad.Quad) {
	if ovf.distRight != math.MinInt32 && ovf.distLeft == math.MinInt32 {
		ovf.parseHorizontalAxis()
	}
	if ovf.distBottom != math.MinInt32 && ovf.distTop == math.MinInt32 {
		ovf.parseVerticalAxis()
	}
	if ovf.distTop < 0 && ovf.distTop != math.MinInt32 {
		ovf.adjustOverlayVerticalOffset(q)
	} else if ovf.distTop < 0 {
		ovf.distTop = 0
	}
	if ovf.distLeft < 0 && ovf.distLeft != math.MinInt32 {
		ovf.adjustOverlayHorizontalOffset(q)
	} else if ovf.distLeft < 0 {
		ovf.distLeft = 0
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

// adjust vertical offset, i.e, removes some lines from top
func (ovf *OverlayFilter) adjustOverlayVerticalOffset(q *quad.Quad) error {
	absOffset := int(math.Abs(float64(ovf.distTop)))
	if err := q.CropTop(absOffset); err != nil {
		return err
	}
	ovf.distTop = 0
	return nil
}

func (ovf *OverlayFilter) adjustOverlayHorizontalOffset(q *quad.Quad) error {
	absOffset := int(math.Abs(float64(ovf.distLeft)))
	if err := q.CropLeft(absOffset); err != nil {
		return err
	}
	ovf.distLeft = 0
	return nil
}
