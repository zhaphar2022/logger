package logger

import "log"

func Info(msg string) {
	log.Printf("INFO - %v", msg)
}

func Warning(msg string) {
	log.Printf("WARN - %v", msg)
}

func Error(msg string) {
	log.Printf("ERROR - %v", msg)
}
