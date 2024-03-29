package logger

import (
	"fmt"
	"log"
	"os"
)

var usage string = `
General Options:
	-h               Display tool usage
	-v               Verbose mode

Input and Output Options (mandatory):
	-i string        Input file
	-o string        Output file

Crop Options:
	-c string        Crop image at given coordinates
										Format: "-c xstart,xend,ystart,yend" or "-c xend,yend"

Image Processing Options:
	-nn bool         Indicate nearest neighbor resize algorithm must be used
	-width int       Output width
	-height int      Output height
	-q int           Quality of the JPEG image (1-100)
	-ssr int         Subsample ratio for images YCbCr
	-fx              Flip x-axis filter
	-fy              Flip y-axis filter
	-gs              Apply grayscale filter
	-pb              Add padding to the background image for the overlay to ensure full visibility. 
	-t               Apply transpose process (rotate 270 degrees and flip Y axis)
	-tl              Rotate 90 degrees counterclockwise
	-tr              Rotate 90 degrees clockwise
	-ov string       Image to overlay onto the input image
	-dt int          (overlay) Distance to the top
	-db int          (overlay) Distance to the bottom
	-dl int          (overlay) Distance to the left
	-dr int          (overlay) Distance to the right
	-l  int          Brightness value. It will be added to each channel for each pixel.
	-b  int          How blurry the image will be. It is the size of the kernel that will be applied.
	-s  int          Sigma value for blur algorithm. It is optional. When not given, it will be half the value of b flag for better result.
`

type Logger struct {
	Warn    bool
	Process bool
	logger  *log.Logger
}

func NewLogger(verbose bool) Logger {
	flags := log.Ldate | log.Ltime // log date and time
	logger := log.New(os.Stdout, "", flags)
	return Logger{
		Warn:    verbose,
		Process: verbose,
		logger:  logger,
	}
}

func (l *Logger) Log(msg string) {
	if l.Process {
		l.logger.Println(msg)
	}
}

func (l *Logger) LogWarnf(format string, msg ...any) {
	if l.Warn {
		l.logger.Printf(format, msg...)
	}
}

func (l *Logger) LogProcessf(format string, msg ...any) {
	if l.Process {
		l.logger.Printf(format, msg...)
	}
}

func (l *Logger) LogWarn(msg string) {
	if l.Warn {
		l.logger.Println(msg)
	}
}

func (l *Logger) LogProcess(msg string) {
	if l.Process {
		l.logger.Println(msg)
	}
}

// logs the message and kills the process with the given status code
func (l *Logger) Fatal(message string, status int) {
	l.logger.Println(message)
	os.Exit(status)
}

func (l *Logger) Usage() {
	fmt.Print(usage)
	os.Exit(0)
}

func Usage() {
	fmt.Print(usage)
	os.Exit(0)
}
