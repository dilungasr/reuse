/*
Copyright © 2022 Dilunga SR <dilungasr@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dilungasr/reuse/errs"
	"github.com/dilungasr/reuse/handlers"
	"github.com/dilungasr/reuse/logger"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "reuse",
	Short: "Automate changes to be made on your current project folder so you can reuse it for the next project",
	Long: `
Reuse provides a simple way to automate changes to be made on your current project folder so you can reuse it for the next project.

- Perform project-wise replacement to any text in your code
- Automate your project setup commands
- Delete files you don't want in the next project
- Specify extensions to modify contents on
- Ignore files you don't want to modify
- Use terminal interactive mode or define everying in config file
- Supports yaml, toml and json... There you have it! 

Don't start from scratch, don't get stressed...
Just reuse it :)

LEARN MORE:
> https://www.github.com/dilungasr/reuse

Much love from Tanzania!
- By Dilunga SR (Sam) <dilungasr@gmail.com>
`,

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

		IsInteractive, err := cmd.Flags().GetBool("interact")
		errs.Check(err)

		IsConfigProvided := use != ""

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
				fmt.Println("Attempting to open in interactive mode...")
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goreuse.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("use", "u", "", `
    Use this to provide a path to the supported configuration file if any. 
	You do not need this in an interactive mode
	`)

	rootCmd.PersistentFlags().BoolP("interact", "i", false, `
	Use this to start the reuse process in an interactive mode
	`)
}
