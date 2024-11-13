package parsing

import (
	"image"
	"os"

	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/chai2010/webp"
)

type WebpParsingStrategy struct {
	logger *logger.Logger
}

func (jps *WebpParsingStrategy) Save(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	err = webp.Encode(fg, img, &webp.Options{Lossless: true})
	if err == nil {
		jps.logger.LogProcess("Image successfully encoded as WEBP")
	}
	return err
}
