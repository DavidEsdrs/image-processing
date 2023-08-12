package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

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
}

// Convert the image into a tensor to further manipulation
func convertIntoTensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	pixels := make([][]color.Color, size.Y)

	for y := 0; y < size.Y; y++ {
		pixels[y] = make([]color.Color, size.X)
		for x := 0; x < size.X; x++ {
			pixels[y][x] = img.At(x, y)
		}
	}

	return pixels
}

func convertIntoImage(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels[0]), len(pixels))
	nImg := image.NewRGBA(rect)
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[0]); x++ {
			p := pixels[y][x]
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	return nImg
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

	return &proc
}

func saveImage(img image.Image, outputPath string) error {
	fg, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fg.Close()
	err = jpeg.Encode(fg, img, nil)
	return err
}

func processImage(img image.Image, file string, outputFolder string, proc processor.Processor) {
	tensor := convertIntoTensor(img)
	iep := proc.Execute(&tensor)
	cImg := convertIntoImage(iep)
	outputPath := outputFolder
	println(outputFolder)
	err := saveImage(cImg, outputPath)
	if err != nil {
		log.Fatal(err.Error())
	}
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
