package parsing

import (
	"image"
	"image/png"
	"os"

	"github.com/DavidEsdrs/image-processing/logger"
)

type PngParsingStrategy struct {
	logger *logger.Logger
}

func (jps *PngParsingStrategy) Save(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	err = png.Encode(fg, img)
	if err == nil {
		jps.logger.LogProcess("Image successfully encoded as PNG")
	}
	return err
}
