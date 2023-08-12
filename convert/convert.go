package convert

import (
	"image"
	"image/color"
)

// Convert the image into a tensor to further manipulation
func ConvertIntoTensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	pixels := make([][]color.Color, size.Y)

	for y := 0; y < size.Y; y++ {
		pixels[y] = make([]color.Color, size.X)
		for x := 0; x < size.X; x++ {
			pixels[y][x] = img.At(x, y)
		}
	}

	return pixels
}

// Convert tensor back to image
func ConvertIntoImage(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewRGBA(rect)
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[0]); x++ {
			q := pixels[y]
			if q == nil {
				continue
			}
			p := pixels[y][x]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
}
