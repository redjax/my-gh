package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchStarredRepos fetches starred repositories from GitHub
func FetchStarredRepos(token string) ([]interface{}, error) {
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
