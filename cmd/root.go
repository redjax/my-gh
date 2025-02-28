package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Root command for the application
var rootCmd = &cobra.Command{
	Use:   "my-gh",
	Short: "My GitHub CLI to fetch starred repositories",
}

// Execute the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Define the root flags
	rootCmd.Flags().String("config-file", "config.json", "Path to the config file")
	rootCmd.Flags().String("gh-api-token", "", "GitHub API token (overrides config file and env)")
	rootCmd.Flags().String("output-file", "starred_repositories.json", "Output file for starred repositories")

	// Set defaults for values if not provided via config, env, or flags
	viper.SetDefault("output_file", "starred_repositories.json")

	// Bind environment variables
	viper.SetEnvPrefix("MYGH") // For example MYGH_GH_TOKEN for the GitHub token
	viper.AutomaticEnv()

	// Add subcommands to root
	rootCmd.AddCommand(getCmd) // add the 'get' command
}
