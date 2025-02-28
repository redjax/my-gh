package utils

import (
	"encoding/json"
	"os"
)

// SaveJSON writes JSON data to a file
func SaveJSON(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
