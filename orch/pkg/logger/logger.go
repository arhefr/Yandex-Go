package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type Log interface {
	Info(...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Fatal(...interface{})
}

type Logger struct {
	Log
}

func NewLogger(file *os.File, level log.Level) *Logger {
	logger := log.New()
	logger.SetOutput(file)
	logger.SetLevel(level)
	return &Logger{logger}
}

func (l *Logger) Info(args ...interface{}) {
	fmt.Println(args...)
	l.Log.Info(args)
}

func (l *Logger) Debug(args ...interface{}) {
	fmt.Println(args...)
	l.Log.Debug(args)
}

func (l *Logger) Warn(args ...interface{}) {
	fmt.Println(args...)
	l.Log.Warn(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	fmt.Println(args...)
	l.Log.Fatal(args)
}
