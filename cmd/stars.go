package cmd

import (
	"fmt"
	"os"

	"redjax/my-gh/internal/utils"

	"github.com/spf13/cobra"
)

var starsCmd = &cobra.Command{
	Use:   "stars",
	Short: "Fetch and save your starred repositories from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		// Load config
		configFile, _ := cmd.Flags().GetString("config-file")
		config, err := utils.LoadConfig(configFile)
		if err != nil {
			fmt.Println("Error loading config:", err)
			os.Exit(1)
		}

		// Fetch starred repositories
		starredRepos, err := utils.FetchStarredRepos(config.GitHubToken)
		if err != nil {
			fmt.Println("Error fetching starred repositories:", err)
			os.Exit(1)
		}

		// Save JSON output
		if err := utils.SaveJSON(config.OutputFile, starredRepos); err != nil {
			fmt.Println("Error saving JSON:", err)
			os.Exit(1)
		}

		fmt.Println("Starred repositories saved to", config.OutputFile)
	},
}

func init() {
	// Define flags for "stars" command
	starsCmd.Flags().String("config-file", "config.json", "Path to the config file")
	starsCmd.Flags().String("gh-api-token", "", "GitHub API token (overrides config file and env)")
	starsCmd.Flags().String("output-file", "starred_repositories.json", "Output file for starred repositories")
}
