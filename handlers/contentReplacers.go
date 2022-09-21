package handlers

import (
	"os"
	"strings"

	"github.com/dilungasr/reuse/models"
)

func ReadFile(path string, info os.FileInfo, level ...int) (content string, ok bool) {
	data, err := os.ReadFile(path)

	if err != nil {
		IndentedOutput(info, "REPLACE FAILED", level...)
		return content, false
	}

	return string(data), true
}

func ReplaceFileContent(content string) (newContent string) {
	newContent = content

	for old, new := range models.ProjectChanges.Rep {
		newContent = strings.ReplaceAll(newContent, old, new)
	}

	return newContent
}

func RewriteFile(filePath, newContent string, info os.FileInfo, level ...int) (ok bool) {
	f, err := os.Create(filePath)
	if err != nil {
		IndentedOutput(info, "REPLACE FAILED", level...)
		return false
	}

	defer f.Close()

	f.WriteString(newContent)
	err = f.Sync()

	if err != nil {
		IndentedOutput(info, "REPLACE FAILED", level...)
		return false
	}

	IndentedOutputNoHyph(info, "✓", level...)
	return true
}
