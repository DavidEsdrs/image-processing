package parsing

import (
	"image"
	"image/png"
	"os"
)

type PngParsingStrategy struct{}

func (jps *PngParsingStrategy) Save(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	err = png.Encode(fg, img)
	return err
}
