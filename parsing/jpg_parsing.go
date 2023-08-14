package parsing

import (
	"image"
	"image/jpeg"
	"os"
)

type JpgParsingStrategy struct{}

func (jps *JpgParsingStrategy) Save(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	err = jpeg.Encode(fg, img, &jpeg.Options{
		Quality: 100,
	})
	return err
}
