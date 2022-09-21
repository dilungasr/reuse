package utils

import (
	"encoding/json"
	"os"

	"github.com/dilungasr/reuse/logger"
	"github.com/dilungasr/reuse/models"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// readConfigContent reads the content of the named file
// and returns them as []byte
func readConfigContent(filePath string) (content []byte) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Fatal(err)
	}

	return content
}

// ReadJSON reads the project changes' file  as JSON
func ReadJSON(filePath string) {
	content := readConfigContent(filePath)

	err := json.Unmarshal(content, &models.ProjectChanges)
	if err != nil {
		logger.Fatal(err)
	}
}

// ReadYAML reads the project changes' file as YAML
func ReadYAML(filePath string) {
	content := readConfigContent(filePath)

	err := yaml.Unmarshal(content, &models.ProjectChanges)
	if err != nil {
		logger.Fatal(err)
	}
}

// ReadTOML reads the project changes' file as TOML
func ReadTOML(filePath string) {
	content := readConfigContent(filePath)

	err := toml.Unmarshal(content, &models.ProjectChanges)
	if err != nil {
		logger.Fatal(err)
	}
}
