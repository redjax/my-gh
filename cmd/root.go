package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "myghstars",
	Short: "Fetch and save your starred repositories from GitHub",
}

// Execute runs the root command
func Execute() {
	// Bind environment variables
	viper.SetEnvPrefix("MYGH")
	viper.AutomaticEnv()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Define flags
	rootCmd.PersistentFlags().String("config-file", "config.json", "Path to the config file")
	rootCmd.PersistentFlags().String("gh-api-token", "", "GitHub API token (overrides config file and env)")
	rootCmd.PersistentFlags().String("output-file", "starred_repositories.json", "Output file for starred repositories")

	// Register subcommands
	rootCmd.AddCommand(runCmd)
}
