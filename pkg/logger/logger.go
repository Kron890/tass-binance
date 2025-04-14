package logger

import (
	"fmt"
	"os"
	"time"
)

var format = "2006-01-02 15:04:05"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Infof(message string, arg ...interface{}) {
	l.logf("INFO", message, arg...)
}

func (l *Logger) Errorf(message string, arg ...interface{}) {
	l.logf("ERROR", message, arg...)

}

func (l *Logger) WarnF(message string, arg ...interface{}) {
	l.logf("WARN", message, arg...)

}

func (l *Logger) logf(level, message string, arg ...interface{}) {
	timestamp := time.Now().Format(format)
	fmt.Fprintf(os.Stdout, "%s [%s]: %s\n", timestamp, level, fmt.Sprintf(message, arg...))
}
