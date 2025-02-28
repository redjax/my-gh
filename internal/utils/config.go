package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config structure
type Config struct {
	GitHubToken string `mapstructure:"gh_token"`
	OutputFile  string `mapstructure:"output_file"`
}

// LoadConfig reads configuration from file, env vars, or flags
func LoadConfig(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file not found. Creating default config:", configFile)

		defaultConfig := Config{
			GitHubToken: "",
			OutputFile:  "starred_repositories.json",
		}

		configData, _ := json.MarshalIndent(defaultConfig, "", "  ")
		os.WriteFile(configFile, configData, 0644)

		fmt.Println("Default config created at", configFile)
		os.Exit(0)
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	config := &Config{
		GitHubToken: viper.GetString("gh_token"),
		OutputFile:  viper.GetString("output_file"),
	}

	return config, nil
}
