package parsing

import (
	"image"
	"image/gif"
	"os"
)

type GifParsingStrategy struct{}

func (jps *GifParsingStrategy) Save(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	err = gif.Encode(fg, img, nil)
	return err
}
