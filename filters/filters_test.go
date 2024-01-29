package filters_test

import (
<<<<<<< HEAD
	"testing"

	"github.com/DavidEsdrs/image-processing/filters"
	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/utils"
)

func TestBlur(t *testing.T) {
	bf, err := filters.NewBlurFilter(logger.Logger{}, 1, 7)
	if err != nil {
		t.FailNow()
	}
	img, err := utils.LoadImage("../images/almoço.png")
	if err != nil {
		t.FailNow()
	}
	tensor := utils.ConvertIntoTensor(img)
	err = bf.Execute(&tensor)
	if err != nil {
		t.FailNow()
	}
=======
	"fmt"
	"image/color"
	"testing"

	"github.com/DavidEsdrs/image-processing/filters"
)

func TestHSL(t *testing.T) {
	red := color.RGBA{24, 98, 118, 255}
	redAsHsl := filters.ColorToHsl(red)
	r, g, b, a := redAsHsl.RGBA()
	res := color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
	fmt.Printf("hsl: %#v\n", redAsHsl)
	fmt.Printf("rgba: RGBA{%v, %v, %v, %v}\n", res.R, res.G, res.B, res.A)
>>>>>>> 4d73d80b2916b9e2ecb7d5b4ef783239d9fc415d
}
