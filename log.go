package main

import (
	"log"
	"os"
)

type Logger struct {
	file *os.File
}

func NewLogger(filePath string) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Logger{file: file}, nil
}

func (l *Logger) Info(message string) {
	log.SetOutput(l.file)
	log.Println("INFO: " + message)
}

func (l *Logger) Error(message string) {
	log.SetOutput(l.file)
	log.Println("ERROR: " + message)
}

func (l *Logger) Close() {
	l.file.Close()
}
