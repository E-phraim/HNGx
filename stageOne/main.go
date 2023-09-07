package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/uno", getInfo)
	http.ListenAndServe(":8080", nil)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	// Get current UTC time
	now := time.Now().UTC()

	// Calculate the time difference in minutes
	timeDifference := now.Minute() - 4 // UTC is 4 minutes behind the server time

	// Check if the time difference is within +/-2 minutes
	if timeDifference < -2 || timeDifference > 2 {
		http.Error(w, "Invalid UTC time", http.StatusBadRequest)
		return
	}

	// GitHub URL of the file being run
	githubFileURL := "https://github.com/username/repo/blob/main/file_name.ext"

	// GitHub URL of the full source code
	githubRepoURL := "https://github.com/username/repo"

	// Create a struct to hold the response data
	response := struct {
		SlackName     string    `json:"slack_name"`
		CurrentDay    string    `json:"current_day"`
		UTCTime       time.Time `json:"utc_time"`
		Track         string    `json:"track"`
		GitHubFileURL string    `json:"github_file_url"`
		GitHubRepoURL string    `json:"github_repo_url"`
		StatusCode    int       `json:"status_code"`
	}{
		SlackName:     slackName,
		CurrentDay:    now.Weekday().String(),
		UTCTime:       now,
		Track:         track,
		GitHubFileURL: githubFileURL,
		GitHubRepoURL: githubRepoURL,
		StatusCode:    http.StatusOK,
	}

	// Encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
