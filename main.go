package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	_ "golang.org/x/image/webp"

	"github.com/DavidEsdrs/image-processing/configs"
	"github.com/DavidEsdrs/image-processing/convert"
	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/parsing"
	"github.com/DavidEsdrs/image-processing/processor"
)

type ProcessResult struct {
	fileName string
	success  bool
}

func processImage(img image.Image, outputPath string, proc processor.Processor, logger logger.Logger) {
	logger.LogProcess("Converting image into tensor")
	tensor := convert.ConvertIntoTensor(img)

	iep := proc.Execute(&tensor)

	context := convert.NewConversionContext(logger)

	conversor, err := context.GetConversor(img, proc.GetColorModel())

	if err != nil {
		log.Fatal(err.Error())
	}

	cImg := conversor.Convert(iep)

	pc := parsing.NewParsingContext(logger)

	config, err := pc.GetConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	logger.LogProcessf("Saving image as %v", outputPath)
	config.Save(cImg, outputPath)
}

func main() {
	var config *configs.Config = configs.GetConfig()
	var verbose bool

	flag.BoolVar(&verbose, "v", false, "Verbose")

	flag.StringVar(&config.Input, "i", "", "Input file")
	flag.StringVar(&config.Output, "o", "", "Output file")

	// filters
	flag.BoolVar(&config.FlipY, "fy", false, "Flip y axis filter")
	flag.BoolVar(&config.FlipX, "fx", false, "Flip x axis filter")
	flag.BoolVar(&config.Transpose, "t", false, "Apply transpose process (rotate 270 degrees and flip Y axis)")
	flag.BoolVar(&config.TurnLeft, "tl", false, "Rotate 90 degrees counterclockwise")
	flag.BoolVar(&config.TurnRight, "tr", false, "Rotate 90 degrees clockwise")
	flag.BoolVar(&config.Grayscale, "gs", false, "Apply grayscale filter")
	flag.Float64Var(&config.NearestNeighbor, "nn", 1.0, "Apply nearest neighbor resize algorithm")
	flag.StringVar(&config.Crop, "c", "", "Crop image at given coordinates. Ex.: \"-c 0,1000,0,200\", xstart,xend,ystart,yend or \"-c 1000,200\", xend,yend (x and y start default to 0)")
	flag.IntVar(&config.Ssr, "ssr", 0, "Subsample ratio for images YCbCr. 444 = 4:4:4, 422 = 4:2:2, 420 = 4:2:0, 440 = 4:4:0, 411 = 4:1:1, 410 = 4:1:0")
	flag.IntVar(&config.Quality, "q", 0, "Quality of the JPEG image. 1-100")

	// overlay
	flag.StringVar(&config.Overlay, "ov", "", "Image to overlay onto the input image")
	flag.IntVar(&config.DistTop, "dt", 0, "Distance to the top")
	flag.IntVar(&config.DistRight, "dr", 0, "Distance to the right")
	flag.IntVar(&config.DistBottom, "db", 0, "Distance to the bottom")
	flag.IntVar(&config.DistLeft, "dl", 0, "Distance to the left")
	flag.BoolVar(&config.Fill, "fill", false, "Should the overlay fill in")

	flag.Parse()

	if config.Input == "" || config.Output == "" {
		flag.Usage()
		log.Fatal("Input and output files are required.")
	}

	results := make([]ProcessResult, 1)

	start := time.Now()

	file := config.Input

	img, err := loadImage(file)

	if err != nil {
		results[0] = ProcessResult{fileName: file, success: false}
		log.Fatalf("error - %v\n", err.Error())
	}

	logger := logger.NewLogger(verbose)

	proc := config.ParseConfig(logger)

	output := config.Output

	// main process
	processImage(img, output, proc, logger)

	duration := time.Since(start)

	logger.LogProcessf("completed: image %v processed - %v\n", file, duration.String())
}

func loadImage(file string) (img image.Image, err error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return
	}
	defer imgFile.Close()
	img, _, err = image.Decode(imgFile)
	return
}
