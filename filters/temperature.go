package filters

import (
	"image/color"
	"math"

	"github.com/DavidEsdrs/image-processing/quad"
	"github.com/DavidEsdrs/image-processing/utils"
)

type TemperatureFilter struct {
	temperature float64
}

func NewTemperatureFilter(temperature float64) *TemperatureFilter {
	return &TemperatureFilter{
		temperature: utils.Clamp(temperature, 1000, 40000),
	}
}

func (tf *TemperatureFilter) Execute(q *quad.Quad) error {
	targetColor := tf.getColorFromTemp()

	q.Apply(func(pixel color.RGBA) color.RGBA {
		newR := uint8(float64(pixel.R) * float64(targetColor.R) / 255)
		newG := uint8(float64(pixel.G) * float64(targetColor.G) / 255)
		newB := uint8(float64(pixel.B) * float64(targetColor.B) / 255)
		return color.RGBA{newR, newG, newB, pixel.A}
	})

	return nil
}

// reference for this amazing algorithm: https://tannerhelland.com/2012/09/18/convert-temperature-rgb-algorithm-code.html
func (tf *TemperatureFilter) getColorFromTemp() color.RGBA {
	var (
		r, g, b float64
	)

	temp := tf.temperature / 100

	// calculate red
	if temp <= 66 {
		r = 255
	} else {
		r = temp - 60
		r = 329.698727446 * (math.Pow(r, -0.1332047592))
		if r < 0 {
			r = 0
		} else if r > 255 {
			r = 255
		}
	}

	// calculate green
	if temp <= 66 {
		g = temp
		g = 99.4708025861*math.Log(g) - 161.1195681661
		if g < 0 {
			g = 0
		} else if g > 255 {
			g = 255
		}
	} else {
		g = temp - 60
		g = 288.1221695283 * (math.Pow(g, -0.0755148492))
		if g < 0 {
			g = 0
		} else if g > 255 {
			g = 255
		}
	}

	// calculate blue
	if temp >= 66 {
		b = 255
	} else {
		if temp <= 19 {
			b = 0
		} else {
			b = temp - 10
			b = 138.5177312231*math.Log(b) - 305.0447927307
			if b < 0 {
				b = 0
			} else if b > 255 {
				b = 255
			}
		}
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}
