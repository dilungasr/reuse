package handlers

import (
	"path/filepath"
	"strings"

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
