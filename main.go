package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/DavidEsdrs/image-processing/configs"
	"github.com/DavidEsdrs/image-processing/convert"
	"github.com/DavidEsdrs/image-processing/parsing"
	"github.com/DavidEsdrs/image-processing/processor"
)

type ProcessResult struct {
	fileName string
	success  bool
}

func processImage(img image.Image, file string, outputPath string, proc processor.Processor) {
	tensor := convert.ConvertIntoTensor(img)

	iep := proc.Execute(&tensor)

	context := convert.NewConversionContext()

	conversor, err := context.GetConversor(img, proc.GetColorModel())

	if err != nil {
		log.Fatal(err.Error())
	}

	cImg := conversor.Convert(iep)

	pc := parsing.NewParsingContext()

	config, err := pc.GetConfig(file)

	if err != nil {
		log.Fatal(err.Error())
	}

	config.Save(cImg, outputPath)
}

func main() {
	var config *configs.Config = configs.GetConfig()

	flag.StringVar(&config.Input, "i", "", "Input file")
	flag.StringVar(&config.Output, "o", "", "Output file")
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

	proc := config.ParseConfig()

	output := config.Output

	// main process
	processImage(img, file, output, proc)

	fmt.Printf("process: image %v processed\n", file)

	duration := time.Since(start)

	fmt.Printf("completed: %v image processed - %v milliseconds\n", 1, duration.Milliseconds())
}

func loadImage(file string) (image.Image, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}
