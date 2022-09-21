package handlers

import (
	"os"
)

// Deleter removes the target file or directory and all of its children
func Deleter(filePath string, info os.FileInfo, level ...int) {
	// if is a directory then remove all
	if info.IsDir() {
		err := os.RemoveAll(filePath)
		if err != nil {
			IndentedOutput(info, "FAILED TO DELETE", level...)
		}
	} else {
		// if is just a file
		err := os.Remove(filePath)
		if err != nil {
			IndentedOutput(info, "FAILED TO DELETE", level...)
		}
	}

	// successul deleted
	IndentedOutput(info, "DELETED", level...)
}
