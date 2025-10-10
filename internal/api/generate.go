package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GenerateRequest represents a text generation request
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// GenerateResponse represents a text generation response
type GenerateResponse struct {
	Response        string `json:"response"`
	TokensGenerated int    `json:"tokens_generated"`
	LatencyMs       int64  `json:"latency_ms"`
}

// Generate sends a text generation request to the daemon
func Generate(model, prompt string) (*GenerateResponse, error) {
	req := GenerateRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	resp, err := http.Post(BaseURL+"/generate", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("post generate: %w", err)
	}
	defer resp.Body.Close()

	rspBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("generate failed: %s", rspBody)
	}

	var out GenerateResponse
	if err := json.Unmarshal(rspBody, &out); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &out, nil
}
