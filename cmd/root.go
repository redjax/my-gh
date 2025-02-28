package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mygh",
	Short: "My GitHub CLI tool",
	Long:  "A command line tool to interact with GitHub repositories, stars, and more.",
}

// Execute runs the root command and initiates subcommands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Register subcommands
	rootCmd.AddCommand(getCmd)
}
