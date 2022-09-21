package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/dilungasr/reuse/logger"
	"github.com/dilungasr/reuse/models"
)

// Run interprets and runs the commands defined in the run field of the config
func Run(projectFolder string) {
	// only inform of executing commands only when there are commands
	// provided
	if len(models.ProjectChanges.Run) > 0 {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("--------Executing setup commands-------")
		fmt.Println("")
	}

	//   iterate the command lines
	for _, commandLine := range models.ProjectChanges.Run {
		//   break the full command in single command, the in the indivial components
		command, args := commandSeparator(commandLine)

		cmd := exec.Command(command, args...)
		// set the working directory of the command
		cmd.Dir = projectFolder

		//all outputs should be printed out to the standard output && all
		// input should be taken from the standard input
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		// run the command and handle the raise error if any
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}

	}
}

func commandSeparator(commandLine string) (command string, args []string) {
	// remove trailing and leading spaces
	commandLine = strings.TrimSpace(commandLine)

	// regular expression for finding quoted  substrings in the command string
	re := regexp.MustCompile(`('[^']*')|("[^"]*")`)

	// stores a slice of maps of placeholder with its value
	phdVals := []map[string]string{}

	i := 0
	phdCmd := re.ReplaceAllStringFunc(commandLine, func(s string) string {
		phd := "$reuse$placeholder$" + strconv.Itoa(i)

		//   record the placeholder value
		phdVals = append(phdVals, map[string]string{
			"phd": phd,
			"val": s[1 : len(s)-1],
		})

		//increment the match count
		i++

		return phd
	})

	// separate the command into args
	cmdArgsRaw := strings.Split(phdCmd, " ")

	for _, cmdArgRaw := range cmdArgsRaw {
		cmdArg := strings.TrimSpace(cmdArgRaw)
		args = append(args, cmdArg)
	}

	// place the values in the placeholders
	for _, phdVal := range phdVals {
		phd := phdVal["phd"]
		val := phdVal["val"]

		//    find the match in the args
		for i, arg := range args {
			if arg == phd {
				args[i] = val
			}
		}
	}
	// handle empty command line
	if len(args) == 0 {
		logger.Fatal("empty command-lines are not allowed")
	}

	// extract command from args and args from args
	command = args[0]

	// if args provided
	if len(args) > 1 {
		args = args[1:]
	}

	return command, args
}
