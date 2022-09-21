/*
Copyright © 2022 Dilunga SR <dilungasr@gmail.com>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/dilungasr/reuse/errs"
	"github.com/dilungasr/reuse/handlers"
	"github.com/dilungasr/reuse/logger"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start to reuse the project in the given path",
	Long: `This command starts the reuse process on the path provided as the first argument.
    The most common way to use this command would be: 
    reuse start ~/myProjects/helloWorld u- @changes/config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		// starting time
		start := time.Now()

		// projectPath defaults to the current directory
		projectPath := "."

		// if project path given then use it
		if len(args) >= 1 && args[0] != "" {
			projectPath = filepath.FromSlash(args[0])
		}

		// look for the flags provided
		use, err := cmd.Flags().GetString("use")
		errs.Check(err)

		interactive, err := cmd.Flags().GetString("interact")
		errs.Check(err)

		IsConfigProvided := use != ""
		IsInteractive := interactive != ""

		// interactively take configurations from the user
		// if it's in interative mode
		if IsInteractive {
			handlers.ReadFromInput()
		} else if IsConfigProvided {
			// read the configuration file if there is a config file to use
			handlers.ReadConfig(use, projectPath)
		} else {
			// search configuration file with default config name 'reuse.supportedExt'
			ok := handlers.FindConfigAndRead(projectPath)

			if !ok {
				fmt.Println("Attempting to open an interactive mode...")
				handlers.ReadFromInput()
			}
		}

		// // start iterating for deleting and replacing

		fmt.Println("")
		fmt.Println("")
		fmt.Println("Starts modifying contents....")
		start2 := time.Now()
		fmt.Println("")

		handlers.Iterate(projectPath)

		// time eplapsed
		logger.Elapsed(start2, "Finished modifying contents in")

		// // finish by running the commands
		handlers.Run(projectPath)
		logger.Elapsed(start, ";) You are almost there in just")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.PersistentFlags().StringP("use", "u", "", `Provides a path to the configuration file if any. 
	You do not need this in interactive mode`,
	)
	startCmd.PersistentFlags().StringP("interact", "i", "", `Start reuse process in interactive mode`)
}
