package logger

import (
	"fmt"
	"strconv"
	"time"
)

// Elapsed logs the elapsed time message
func Elapsed(start time.Time, msg string) {
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("")
	fmt.Println(msg, strconv.Itoa(int(elapsed))+"ms")
}
