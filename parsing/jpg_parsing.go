package parsing

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/DavidEsdrs/image-processing/configs"
	"github.com/DavidEsdrs/image-processing/logger"
)

type JpgParsingStrategy struct {
	logger *logger.Logger
}

func (jps *JpgParsingStrategy) Save(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	quality := configs.GetConfig().Quality
	if quality > 0 {
		err = jpeg.Encode(fg, img, &jpeg.Options{
			Quality: quality,
		})
	} else {
		err = jpeg.Encode(fg, img, nil)
	}
	if err == nil {
		jps.logger.LogProcess("Image successfully encoded as JPEG")
	}
	return err
}
