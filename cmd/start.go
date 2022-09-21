/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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

		// check if use flag provided
		use, err := cmd.Flags().GetString("use")
		errs.Check(err)

		IsConfigProvided := use != ""

		// read the configuration file  is there is a config file to use
		if IsConfigProvided {
			handlers.ReadConfig(use, projectPath)
		} else {
			// just go for the console prompts
			handlers.ReadFromInput()
		}

		// // start iterating for deleting and replacing
		fmt.Println("")
		fmt.Println("")
		fmt.Println("--------UPDATING PROJECT CONTENTS-------")
		fmt.Println("")

		handlers.Iterate(projectPath)

		// time eplapsed
		logger.Elapsed(start, "FINISHED UPDATING CONTENTS IN")

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
	You do not need this if you want to set your configurations directly on the terminal`,
	)
}
