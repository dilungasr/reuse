package handlers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dilungasr/reuse/errs"
	"github.com/dilungasr/reuse/models"

	"strings"
)

// Iterator iterates the project folder to find the files or directories subjected to
// changes
func Iterate(filePath string, level ...int) {

	entries, err := os.ReadDir(filePath)
	if err != nil {
		log.Panicln(err)
	}

	if level == nil {
		level = []int{1}
	}

	for _, entry := range entries {
		info, err := entry.Info()

		errs.Check(err)

		fileName := info.Name()

		// file or folder path
		path := filepath.Join(filePath, fileName)

		// files to skip
		if IsSubjecTo(models.ProjectChanges.Ignore, path) {
			continue
		}

		// file to delete
		if IsSubjecTo(models.ProjectChanges.Del, path) {
			Deleter(path, info, level...)
			continue
		}

		// get deeper if it's a directory
		if info.IsDir() {
			IndentedOutput(info, "", level...)
			Iterate(path, level[0]+1)
		} else {
			// get the file extension
			nameParts := strings.Split(fileName, ".")
			if len(nameParts) == 1 {
				continue
			}

			ext := nameParts[1]
			// fmt.Println(ext)

			// only read the file if it's a required file extension
			if IsToRead(ext) {
				content, ok := ReadFile(path, info, level...)
				if !ok {
					continue
				}

				//  make changes to the file
				newContent := ReplaceFileContent(content)

				// write the new content to the file
				RewriteFile(path, newContent, info, level...)
			} else {
				continue
			}
		}

	}
}
