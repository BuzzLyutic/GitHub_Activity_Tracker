package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	//"github.com/urfave/cli/v2"
)

const GITHUBAPIURL = "https://api.github.com/users/%s/events"

type ApiResponse struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	Repo      struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func fetchGitHubActivity(username string) ([]ApiResponse, error) {

	url := fmt.Sprintf(GITHUBAPIURL, username)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch data from GitHub: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("GitHub API returned status %d: %s", response.StatusCode, body)
	}

	var results []ApiResponse

	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return results, nil

}

func main() {

	fmt.Println("The GitHub Activity Tracker Program")

	// Get username
	var username string
	fmt.Print("Please enter your GitHub username: ")
	if _, err := fmt.Scan(&username); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Fetch GitHub activity
	results, err := fetchGitHubActivity(username)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Display results
	if len(results) == 0 {
		fmt.Println("No recent activity found for this user.")
		return
	}

	fmt.Println("\nRecent GitHub Activity:")
	for _, result := range results {
		fmt.Printf("  - Type: %s | Created at: %s | Repo: %s\n",
			result.Type, result.CreatedAt, result.Repo.Name)
	}
}
