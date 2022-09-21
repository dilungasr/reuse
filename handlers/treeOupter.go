package handlers

import (
	"fmt"
	"os"
	"strings"
)

var indentSpace = "   "

func IndentedOutput(info os.FileInfo, action string, level ...int) {

	// format the aciton if provided
	if action != "" {
		action = "- " + strings.ToUpper(action)
	}

	// give output by indenting depending on the folder tree level
	if level[0] == 1 {
		fmt.Println(info.Name(), action)
	} else {
		indent := strings.Repeat(indentSpace, level[0]-1)
		fmt.Println(indent, info.Name(), action)
	}
}
func IndentedOutputNoHyph(info os.FileInfo, action string, level ...int) {

	// format the aciton if provided
	if action != "" {
		action = strings.ToUpper(action)
	}

	// give output by indenting depending on the folder tree level
	if level[0] == 1 {
		fmt.Println(info.Name(), action)
	} else {
		indent := strings.Repeat(indentSpace, level[0]-1)
		fmt.Println(indent, info.Name(), action)
	}
}
