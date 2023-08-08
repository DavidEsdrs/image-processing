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
	"path/filepath"
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
	NearestNeighbor float64
}

// Convert the image into a tensor to further manipulation
func convertIntoTensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	var pixels [][]color.Color
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}
	return pixels
}

func convertIntoImage(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)
	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
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
	outputPath := filepath.Join(outputFolder, file)
	saveImage(cImg, outputPath)
}

func main() {
	var config Config

	flag.StringVar(&config.Input, "-i", "", "Input file")
	flag.StringVar(&config.Output, "-o", "", "Output file")
	flag.BoolVar(&config.FlipY, "-fy", false, "Flip y axis filter")
	flag.BoolVar(&config.FlipX, "-fx", false, "Flip x axis filter")
	flag.BoolVar(&config.Transpose, "-t", false, "Apply transpose process (rotate 270 degrees and flip Y axis)")
	flag.Float64Var(&config.NearestNeighbor, "-nn", 1.0, "Apply nearest neighbor resize algorithm")

	results := make([]ProcessResult, 1)

	start := time.Now()

	file := config.Input

	img, err := loadImage(file)

	if err != nil {
		results[0] = ProcessResult{fileName: file, success: false}
		log.Fatalf("error - %v\n", err.Error())
	}

	proc := parseConfig(config)

	// main process
	processImage(img, file, "assets", proc)

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
