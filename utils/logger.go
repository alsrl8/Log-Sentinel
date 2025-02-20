package utils

import (
	"log"
	"os"
	"sync"
)

var (
	logger *Logger
	once   sync.Once
)

type Logger struct {
	*log.Logger
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.Printf("[Info] "+format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.Printf("[Debug] "+format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.Printf("[Warn] "+format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.Printf("[Error] "+format, v...)
}

func GetLogger() *Logger {
	once.Do(func() {
		logPath := "./logs/app.log"
		logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		logger = &Logger{log.New(logFile, "", log.LstdFlags)}

	})
	return logger
}
