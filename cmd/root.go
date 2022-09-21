/*
Copyright © 2022 Dilunga SR <dilungasr@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goreuse",
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
