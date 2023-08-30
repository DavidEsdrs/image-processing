package utils

import "image/color"

func ColorModelString(model color.Model) string {
	switch model {
	case color.Alpha16Model:
		return "Alpha16Model"
	case color.AlphaModel:
		return "AlphaModel"
	case color.CMYKModel:
		return "CMYKModel"
	case color.Gray16Model:
		return "Gray16Model"
	case color.GrayModel:
		return "GrayModel"
	case color.NRGBA64Model:
		return "NRGBA64Model"
	case color.NRGBAModel:
		return "NRGBAModel"
	case color.RGBA64Model:
		return "RGBA64Model"
	case color.RGBAModel:
		return "RGBAModel"
	case color.YCbCrModel:
		return "YCbCrModel"
	default:
		return "unknown color model"
	}
}
