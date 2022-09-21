package handlers

import (
	"path/filepath"

	"github.com/dilungasr/reuse/models"
)

// IsSubjectTo checks if the given file is subjected to specific action
// in the changes config
//
// e.g is subjected to skip or delete action
func IsSubjecTo(subjectedFiles []string, filePath string) bool {
	// get the file name from the file path
	_, file := filepath.Split(filePath)

	//compare to see is subjected
	for _, fileName := range subjectedFiles {
		if fileName == file || fileName == filePath {
			return true
		}
	}

	return false
}

// IsToRead checks if ext is provided for modification
func IsToRead(ext string) bool {
	for _, toReadExt := range models.ProjectChanges.Ext {
		if ext == toReadExt {
			return true
		}
	}

	return false
}
