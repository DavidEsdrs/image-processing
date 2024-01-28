package hsl

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/utils"
)

type HSL struct {
	H float32
	S float32
	L float32
}

func (hsl HSL) RGBA() (r, g, b, a uint32) {
	if hsl.S == 0 {
		r = uint32(0xffff * hsl.L)
		g = uint32(0xffff * hsl.L)
		b = uint32(0xffff * hsl.L)
		a = 0xffff
		return
	}

	var (
		temp1 float32
		temp2 float32
		tempR float32
		tempG float32
		tempB float32
	)

	if hsl.L < 0.5 {
		temp1 = hsl.L * (1 + hsl.S)
	} else {
		temp1 = hsl.L + hsl.S - hsl.L*hsl.S
	}

	temp2 = 2*hsl.L - temp1

	fixedH := float32(hsl.H) / 360

	tempR = fixedH + 1.0/3

	if tempR > 1 {
		tempR -= 1
	}

	tempG = fixedH

	tempB = fixedH - 1.0/3
	if tempB < 0 {
		tempB += 1
	}

	r = uint32(hueToRGB(tempR, temp1, temp2) * 0xffff)
	g = uint32(hueToRGB(tempG, temp1, temp2) * 0xffff)
	b = uint32(hueToRGB(tempB, temp1, temp2) * 0xffff)
	a = 0xffff // we'll treat it as a fully opaque pixel

	return
}

func ColorToHsl(color color.Color) HSL {
	var (
		H float32
		S float32
		L float32
	)

	R, G, B, _ := color.RGBA()

	r := float32(R) / 0xffff
	g := float32(G) / 0xffff
	b := float32(B) / 0xffff

	min := utils.Min(r, g, b)
	max := utils.Max(r, g, b)
	c := max - min

	L = float32(min+max) / 2

	if min == max {
		H = 0
		S = 0
	} else {
		if L <= 0.5 {
			S = float32(c) / float32(max+min)
		} else {
			S = float32(c) / float32(2-c)
		}

		switch max {
		case r:
			H = (g - b) / c
		case g:
			H = 2 + (b-r)/c
		case b:
			H = 4 + (r-g)/c
		}

		H *= 60

		if H < 0 {
			H += 360
		}
	}

	return HSL{H, S, L}
}

func hueToRGB(channel, temp1, temp2 float32) float32 {
	if 6*channel < 1 {
		return temp2 + (temp1-temp2)*6*channel
	} else if 2*channel < 1 {
		return temp1
	} else if 3*channel < 2 {
		return temp2 + (temp1-temp2)*(2.0/3-channel)*6
	}

	return temp2
}
