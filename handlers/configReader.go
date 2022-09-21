package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dilungasr/reuse/errs"
	"github.com/dilungasr/reuse/logger"
	"github.com/dilungasr/reuse/utils"
)

func ReadConfig(filePath, projectFolder string) {
	// interpret the @ symbol as the root of the project folder
	if filePath[0] == '@' {
		filePath = strings.Replace(filePath, "@", "", 1)
		filePath = filepath.Join(projectFolder, filePath)
	}

	//read content as defined by the extension
	ext := filepath.Ext(filePath)

	switch ext {
	case ".toml":
		utils.ReadTOML(filePath)
	case ".yaml", ".yml":
		utils.ReadYAML(filePath)
	case ".json":
		utils.ReadJSON(filePath)
	default:
		_, fileName := filepath.Split(filePath)
		logger.Fatal("Unsupported file type (" + fileName + ")")
	}
}

func FindConfigAndRead(projectFolder string) {
	entries, err := os.ReadDir(projectFolder)
	errs.Check(err)

	// find the first found config with default name
	for _, entry := range entries {

		// skip directories
		if entry.IsDir() {
			continue
		}

		// check for default names matching on the supported config extensions
		name := entry.Name()

		if name == "reuse.toml" || name == "reuse.yaml" ||
			name == "reuse.yml" || name == "reuse.json" {
			//    join the fileName with the project folder to create
			// file path
			filePath := filepath.Join(projectFolder, name)
			ReadConfig(filePath, projectFolder)
			break
		}
	}
}
