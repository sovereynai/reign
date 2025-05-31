package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type JobRequest struct {
	Image      string `json:"image"`
	Model      string `json:"model"`
	Redundancy int    `json:"redundancy"`
}

type JobResponse struct {
	JobID string `json:"job_id"`
	Hash  string `json:"proof_hash"`
}

func SubmitJob(imagePath, model string, redundancy int) (*JobResponse, error) {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("read image: %w", err)
	}

	req := JobRequest{
		Image:      string(data),
		Model:      model,
		Redundancy: redundancy,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal job: %w", err)
	}

	resp, err := http.Post("http://localhost:8080/jobs", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("post job: %w", err)
	}
	defer resp.Body.Close()

	rspBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("submit failed: %s", rspBody)
	}

	var out JobResponse
	if err := json.Unmarshal(rspBody, &out); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &out, nil
}
