package logger

import (
	"log"
	"os"
)

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
