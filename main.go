package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Config represents the structure of the config.json file
type Config struct {
	GitHubToken string `json:"gh_token"`
	OutputFile  string `json:"output_file"`
}

func main() {
	// Load config from file
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	if config.GitHubToken == "" {
		fmt.Println("GitHub API token is missing in config.json")
		return
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Define the request URL and headers
	url := "https://api.github.com/user/starred"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	req.Header.Set("User-Agent", "mygh-go")
	req.Header.Set("Authorization", "Bearer "+config.GitHubToken)
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check if the request was successful (status code 200-299)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
		return
	}

	// Decode the JSON response into a Go structure
	var starredRepositories []interface{}
	err = json.Unmarshal(body, &starredRepositories)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Marshal the JSON with indentation
	indentedJSON, err := json.MarshalIndent(starredRepositories, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Save the indented JSON to a file
	err = os.WriteFile(config.OutputFile, indentedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("Response saved to: " + config.OutputFile)
}

// loadConfig reads the config.json file and parses it into a Config struct
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		fmt.Println("Config file not found. Creating a new one...")
		defaultConfig := Config{GitHubToken: ""}
		saveConfig(filename, &defaultConfig)
		fmt.Println("Please update config.json with your GitHub API token.")
		return nil, fmt.Errorf("config file created, please update it with your API token")
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// saveConfig writes the given Config struct to a JSON file
func saveConfig(filename string, config *Config) error {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, configJSON, 0644)
}
