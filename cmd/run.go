package cmd

import (
	"fmt"
	"os"

	"redjax/my-gh/internal/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Fetch and save your starred repositories",
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config-file")

		// Load Config
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
