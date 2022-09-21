package models

import (
	"fmt"
	"strings"
)

var ProjectChanges Changes = Changes{}

// Changes represents the changes to be made to the project
type Changes struct {
	Rep    map[string]string
	Del    []string
	Run    []string
	Ext    []string
	Ignore []string
}

// AddReps adds replacements provided from the config file to the project changes
func (changes *Changes) AddReps(reps map[string]string) {
	changes.Rep = reps
}

// AddRepsFromInput interprets and adds the interpreted replacements to the project changes
func (changes *Changes) AddRepsFromInput(rawReps []string) {
	reps := make(map[string]string)

	for _, rawRep := range rawReps {
		rawRepParts := strings.Split(rawRep, ":")

		//if invalid format
		len := len(rawRepParts)
		if len != 2 {
			fmt.Println("INVALID FORMAT -> ", rawRep)
			continue
		}

		old := strings.TrimSpace(rawRepParts[0])
		new := strings.TrimSpace(rawRepParts[1])

		reps[old] = new
	}

	changes.Rep = reps
}

func (changes *Changes) AddDel(dels []string) {
	changes.Del = dels
}
func (changes *Changes) AddRun(commands []string) {
	changes.Run = commands
}
func (changes *Changes) AddExts(exts []string) {
	changes.Ext = exts
}
func (changes *Changes) AddIgnore(ignore []string) {
	changes.Ignore = ignore
}
