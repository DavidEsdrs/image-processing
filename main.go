package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"time"

	_ "golang.org/x/image/webp"

	"github.com/DavidEsdrs/image-processing/configs"
	"github.com/DavidEsdrs/image-processing/convert"
	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/parsing"
	"github.com/DavidEsdrs/image-processing/processor"
	"github.com/DavidEsdrs/image-processing/utils"
)

// driver code
// Note that application starts and end here - any errors through all the
// application converges to the errors got there on the "if err != nil"
// statements
func main() {
	var config *configs.Config = configs.GetConfig()
	var verbose bool
	var help bool

	setFlags(config, &verbose, &help)

	flag.Parse()

	if help {
		logger.Usage()
		os.Exit(0)
	}

	logger := logger.NewLogger(verbose)

	if config.Input == "" || config.Output == "" {
		logger.Fatal("input and output files are required.", 2)
	}

	start := time.Now()

	img, err := utils.LoadImage(config.Input)

	if err != nil {
		log.Fatalf("error while loading input file - %v\n", err.Error())
	}

	proc, err := config.ParseConfig(logger, img)

	if err != nil {
		log.Fatal(err.Error())
	}

	// main process
	err = processImage(img, config.Output, proc, &logger)

	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(start)

	logger.LogProcessf("completed: image %v processed - output image: %v - %v\n", config.Input, config.Output, duration.String())
}

func processImage(img image.Image, outputPath string, proc *processor.Invoker, logger *logger.Logger) error {
	logger.LogProcess("Converting image into tensor")
	tensor := utils.ConvertIntoTensor(img)

	iep, err := proc.Invoke(&tensor)

	if err != nil {
		return err
	}

	context := convert.NewConversionContext(logger)

	conversor, err := context.GetConversor(img, proc.GetColorModel())

	if err != nil {
		return err
	}

	cImg := conversor.Convert(*iep)

	pc := parsing.NewParsingContext(logger)

	config, err := pc.GetConfig()

	if err != nil {
		return err
	}

	logger.LogProcessf("Saving image as %v", outputPath)
	err = config.Save(cImg, outputPath)
	return err
}

// set cli flags
//
// TODO: Avoid using pointer to primitive args
func setFlags(config *configs.Config, verbose *bool, help *bool) {
	flag.BoolVar(help, "h", false, "Print tool usage")
	flag.BoolVar(verbose, "v", false, "Verbose")

	flag.StringVar(&config.Input, "i", "", "Input file")
	flag.StringVar(&config.Output, "o", "", "Output file")

	// filters
	flag.BoolVar(&config.FlipY, "fy", false, "Flip y axis filter")
	flag.BoolVar(&config.FlipX, "fx", false, "Flip x axis filter")
	flag.BoolVar(&config.Transpose, "t", false, "Apply transpose process (rotate 270 degrees and flip Y axis)")
	flag.BoolVar(&config.TurnLeft, "tl", false, "Rotate 90 degrees counterclockwise")
	flag.BoolVar(&config.TurnRight, "tr", false, "Rotate 90 degrees clockwise")
	flag.BoolVar(&config.Grayscale, "gs", false, "Apply grayscale filter")
	flag.StringVar(&config.Crop, "c", "", "Crop image at given coordinates. Ex.: \"-c 0,1000,0,200\", xstart,xend,ystart,yend or \"-c 1000,200\", xend,yend (x and y start default to 0)")
	flag.IntVar(&config.Ssr, "ssr", 0, "Subsample ratio for images YCbCr. 444 = 4:4:4, 422 = 4:2:2, 420 = 4:2:0, 440 = 4:4:0, 411 = 4:1:1, 410 = 4:1:0")
	flag.IntVar(&config.Quality, "q", 0, "Quality of the JPEG image. 1-100")

	// Resize
	flag.BoolVar(&config.NearestNeighbor, "nn", false, "Apply nearest neighbor resize algorithm")
	flag.IntVar(&config.Width, "width", math.MaxInt32, "Width")
	flag.IntVar(&config.Height, "height", math.MaxInt32, "Height")
	flag.Float64Var(&config.Factor, "f", 1, "Scale factor")

	// overlay
	flag.StringVar(&config.Overlay, "ov", "", "Image to overlay onto the input image")
	flag.IntVar(&config.DistTop, "dt", math.MinInt32, "Distance to the top")
	flag.IntVar(&config.DistRight, "dr", math.MinInt32, "Distance to the right")
	flag.IntVar(&config.DistBottom, "db", math.MinInt32, "Distance to the bottom")
	flag.IntVar(&config.DistLeft, "dl", math.MinInt32, "Distance to the left")
	flag.BoolVar(&config.Fill, "fill", false, "Should the overlay fill in")
}
