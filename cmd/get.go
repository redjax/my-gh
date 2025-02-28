package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"redjax/my-gh/internal/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Define the 'get' subcommand
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch information from GitHub",
	Run:   getStarredRepos, // The logic to fetch starred repos will go here
}

func init() {
	// Optionally, you can add specific flags for the 'get' command here
	// getCmd.Flags().String("some-flag", "default", "Description of flag")
}

// getStarredRepos is the function to run when 'my-gh get' is executed
func getStarredRepos(cmd *cobra.Command, args []string) {
	// Load config file path from the CLI argument
	configFile, _ := cmd.Flags().GetString("config-file")

	// If the config file doesn't exist, check environment variables or CLI args
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		ghToken := viper.GetString("gh_token")
		outputFile := viper.GetString("output_file")

		// If the necessary values are not found, create the config file
		if ghToken == "" && outputFile == "" {
			fmt.Println("Config file not found and no values found in environment or CLI. Creating default config:", configFile)

			defaultConfig := utils.Config{
				GitHubToken: "",
				OutputFile:  "starred_repositories.json",
			}

			// Marshal and create the config file
			configData, _ := json.MarshalIndent(defaultConfig, "", "  ")
			os.WriteFile(configFile, configData, 0644)

			fmt.Println("Default config created at path config.json")
			os.Exit(0) // Exit after creating the config
		}
	}

	// Load config if it exists
	if _, err := os.Stat(configFile); err == nil {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file:", err)
		} else {
			fmt.Println("Loaded config from:", viper.ConfigFileUsed())
		}
	}

	// Bind CLI flags *after* reading config
	viper.BindPFlag("gh_token", cmd.Flags().Lookup("gh-api-token"))
	viper.BindPFlag("output_file", cmd.Flags().Lookup("output-file"))

	// Load values from config, environment, and CLI args
	config := utils.Config{
		GitHubToken: viper.GetString("gh_token"), // From environment, CLI, or config
		OutputFile:  viper.GetString("output_file"),
	}

	// Validate GitHub token
	if config.GitHubToken == "" {
		fmt.Println("GitHub API token is missing. Provide via CLI flag, env variable, or config file.")
		return
	}

	// Fetch starred repositories
	starredRepos, err := utils.FetchStarredRepos(config.GitHubToken)
	if err != nil {
		fmt.Println("Error fetching starred repositories:", err)
		return
	}

	// Save to JSON file
	if err := utils.SaveJSON(config.OutputFile, starredRepos); err != nil {
		fmt.Println("Error saving JSON to file:", err)
		return
	}

	fmt.Println("Starred repositories saved to", config.OutputFile)
}
