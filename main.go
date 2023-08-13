package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/DavidEsdrs/image-processing/convert"
	"github.com/DavidEsdrs/image-processing/parsing"
	"github.com/DavidEsdrs/image-processing/processor"
)

type ProcessResult struct {
	fileName string
	success  bool
}

type Config struct {
	Input           string
	Output          string
	FlipY           bool
	FlipX           bool
	Transpose       bool
	Grayscale       bool
	TurnLeft        bool
	TurnRight       bool
	NearestNeighbor float64
	Crop            string
}

func parseConfig(config Config) processor.Processor {
	proc := processor.ImageProcessor{}

	if config.Transpose {
		proc.Transpose()
	}
	if config.FlipY {
		proc.FlipY()
	}
	if config.FlipX {
		proc.FlipX()
	}
	if config.NearestNeighbor != 1.0 {
		proc.NearestNeighbor(float32(config.NearestNeighbor))
	}
	if config.Grayscale {
		proc.BlackAndWhite()
	}
	if config.TurnLeft {
		proc.TurnLeft()
	}
	if config.TurnRight {
		proc.TurnRight()
	}
	if config.Crop != "" {
		str := strings.Split(config.Crop, ",")

		var xstart int
		var xend int
		var ystart int
		var yend int

		if len(str) == 4 {
			xstart, _ = strconv.Atoi(str[0])
			xend, _ = strconv.Atoi(str[1])
			ystart, _ = strconv.Atoi(str[2])
			yend, _ = strconv.Atoi(str[3])
		} else {
			xend, _ = strconv.Atoi(str[0])
			yend, _ = strconv.Atoi(str[1])
		}

		proc.Crop(xstart, xend, ystart, yend)
	}

	return &proc
}

func processImage(img image.Image, file string, outputPath string, proc processor.Processor) {
	tensor := convert.ConvertIntoTensor(img)
	iep := proc.Execute(&tensor)

	context := convert.NewConversionContext()

	conversor, err := context.GetConversor(file)

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
	var config Config

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

	proc := parseConfig(config)

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
