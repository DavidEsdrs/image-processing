package logger

import "fmt"

type Logger struct {
	Warn    bool
	Process bool
}

func (l *Logger) LogWarn(format, msg string) {
	if l.Warn {
		fmt.Printf(format, msg)
	}
}

func (l *Logger) LogProcess(format, msg string) {
	if l.Process {
		fmt.Printf(format, msg)
	}
}
