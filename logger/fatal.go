package logger

import (
	"log"
)

// LogFatal creates a REUSE error msg and calls log.Fatal()
func Fatal(msg ...any) {
	logMsg := []any{"REUSE: "}
	logMsg = append(logMsg, msg...)
	log.Fatal(logMsg...)
}

// Println creates a REUSE error msg and calls log.Println()
func Println(msg ...any) {
	logMsg := []any{"REUSE: "}
	logMsg = append(logMsg, msg...)
	log.Println(logMsg...)
}
