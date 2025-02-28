package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config structure
type Config struct {
	GitHubToken string `mapstructure:"gh_token"`
	OutputFile  string `mapstructure:"output_file"`
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myghstars",
		Short: "Fetch and save your starred repositories from GitHub",
		Run:   runApp,
	}

	// Define flags
	rootCmd.Flags().String("config-file", "config.json", "Path to the config file")
	rootCmd.Flags().String("gh-token", "", "GitHub API token (overrides config file and env)")
	rootCmd.Flags().String("output-file", "starred_repositories.json", "Output file for starred repositories")

	// Bind flags to Viper before execution
	viper.BindPFlag("gh_token", rootCmd.Flags().Lookup("gh-token"))
	viper.BindPFlag("output_file", rootCmd.Flags().Lookup("output-file"))

	// Bind environment variables
	viper.SetEnvPrefix("MYGH") // Environment variables like MYGH_GH_TOKEN
	viper.AutomaticEnv()

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func runApp(cmd *cobra.Command, args []string) {
	configFile, _ := cmd.Flags().GetString("config-file")

	// Ensure config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file not found. Creating default config:", configFile)

		defaultConfig := Config{
			GitHubToken: "",
			OutputFile:  "starred_repositories.json",
		}

		configData, _ := json.MarshalIndent(defaultConfig, "", "  ")
		os.WriteFile(configFile, configData, 0644)

		fmt.Println("Default config created. Please update it and rerun the command.")
		return // Instead of os.Exit(0)
	}

	// Load the config file
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	} else {
		fmt.Println("Loaded config from:", viper.ConfigFileUsed())
	}

	// Load values from config, environment variables, and CLI flags
	config := Config{
		GitHubToken: viper.GetString("gh_token"),
		OutputFile:  viper.GetString("output_file"),
	}

	// Validate GitHub token
	if config.GitHubToken == "" {
		fmt.Println("GitHub API token is missing. Provide it via CLI flag, environment variable, or config file.")
		return
	}

	// Fetch starred repositories
	starredRepos, err := fetchStarredRepos(config.GitHubToken)
	if err != nil {
		fmt.Println("Error fetching starred repositories:", err)
		return
	}

	// Save to JSON file
	if err := saveJSON(config.OutputFile, starredRepos); err != nil {
		fmt.Println("Error saving JSON:", err)
		return
	}

	fmt.Println("Starred repositories saved to", config.OutputFile)
}

// fetchStarredRepos fetches starred repositories from GitHub
func fetchStarredRepos(token string) ([]interface{}, error) {
	client := &http.Client{}
	url := "https://api.github.com/user/starred"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "mygh-go")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle GitHub API rate limiting
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("GitHub API rate limit exceeded. Try again later")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	var starredRepositories []interface{}
	if err := json.Unmarshal(body, &starredRepositories); err != nil {
		return nil, err
	}

	return starredRepositories, nil
}

// saveJSON writes JSON data to a file
func saveJSON(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
