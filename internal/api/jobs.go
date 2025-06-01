package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Job struct {
	ID     string `json:"id"`
	Model  string `json:"model"`
	Status string `json:"status"`
}

func FetchJobs() ([]Job, error) {
	resp, err := http.Get("http://localhost:8080/jobs")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jobs: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var jobs []Job
	if err := json.Unmarshal(body, &jobs); err != nil {
		return nil, fmt.Errorf("failed to parse jobs: %w", err)
	}

	return jobs, nil
}
