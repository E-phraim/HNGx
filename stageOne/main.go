package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/uno", getInfo)
	fmt.Println("Server Running...")

	http.ListenAndServe(":8080", nil)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")


	utcLocation, err := time.LoadLocation("UTC")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	now := time.Now().In(utcLocation)

	timeDifference := now.Minute() - 4

	if timeDifference < -2 || timeDifference > 2 {
		http.Error(w, "Invalid UTC time", http.StatusBadRequest)
		return
	}

	githubFileURL := "https://github.com/E-phraim/HNGx/blob/main/stageOne/main.go"

	githubRepoURL := "https://github.com/E-phraim/HNGx/tree/main/stageOne"

	response := struct {
		SlackName     string    `json:"kodeforce"`
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
