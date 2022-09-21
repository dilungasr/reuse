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
	Short: "Replace module names and everything you need to reuse your project",
	Long: `Go Reuse provides a simple way to reuse your project to create a new project with 
	some little changes.

- Replacing your module package name to all files in your project so you don't
mess with manual work of updating the imports to refer to the new package
name. 
- Global replacement to any text in your codes
- You can exclude files 
- You can specify the target extensions

Don't start from scratch, don't get stressed...
Just reuse it :)
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
	rootCmd.PersistentFlags().StringP("use", "u", "", `Provides a path to the configuration file if any. 
	You do not need this in interactive mode`,
	)
	rootCmd.PersistentFlags().BoolP("interact", "i", false, `Start reuse process in interactive mode`)
}
