// this file is responsible to transform "DistBottom" (distance to bottom) and
// "DistRight" (distance to right) into "DistTop" and "DistLeft", to allow that
// the image processors just need to sum the DistTop and DistLeft in the indexes

package configs

import (
	"image/color"
	"math"
)

func (config *Config) ParseOverlayConfigs(tensor *[][]color.Color) {
	if (config.DistBottom != -1 || config.DistBottom != config.DistTop) && config.DistTop == 0 {
		parseHorizontalAxis(config)
	}
	if (config.DistRight != -1 || config.DistRight != config.DistLeft) && config.DistLeft == 0 {
		parseVerticalAxis(config)
	}
	if config.DistTop < 0 {
		adjustOverlayVerticalOffset(config, tensor)
	}
	if config.DistLeft < 0 {
		adjustOverlayHorizontalOffset(config, tensor)
	}
}

func parseVerticalAxis(cfg *Config) {
	ovPlustDist := cfg.DistRight + cfg.overlayRect.Max.X
	distToLeft := cfg.backgroundRect.Max.X - ovPlustDist
	cfg.DistLeft = distToLeft
}

func parseHorizontalAxis(cfg *Config) {
	ovPlustDist := cfg.DistBottom + cfg.overlayRect.Max.Y
	distToTop := cfg.backgroundRect.Max.Y - ovPlustDist
	cfg.DistTop = distToTop
}

func adjustOverlayVerticalOffset(config *Config, tensor *[][]color.Color) {
	absOffset := int(math.Abs(float64(config.DistTop)))
	pTensor := (*tensor)[absOffset:]
	*tensor = pTensor
	config.DistTop = 0
}

func adjustOverlayHorizontalOffset(config *Config, tensor *[][]color.Color) {
	absOffset := int(math.Abs(float64(config.DistLeft)))
	pTensor := slice2D(*tensor, 0, absOffset)
	*tensor = pTensor
	config.DistLeft = 0
}

func slice2D(slice [][]color.Color, startRow int, startCol int) [][]color.Color {
	result := make([][]color.Color, len(slice)-startRow)
	for i := startRow; i < len(slice); i++ {
		result[i-startRow] = slice[i][startCol:]
	}
	return result
}
