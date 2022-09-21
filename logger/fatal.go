package logger

import (
	"fmt"
	"log"
	"strconv"
	"time"
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

// Elapsed logs the elapsed time message
func Elapsed(start time.Time, msg string) {
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("")
	fmt.Println(msg, strconv.Itoa(int(elapsed))+"ms")
}
