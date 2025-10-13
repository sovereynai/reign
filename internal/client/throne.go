package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ThroneClient communicates with the throne daemon
type ThroneClient struct {
	BaseURL string
	client  *http.Client
}

// NewThroneClient creates a new throne API client
func NewThroneClient(baseURL string) *ThroneClient {
	return &ThroneClient{
		BaseURL: baseURL,
		client:  &http.Client{Timeout: 60 * time.Second},
	}
}

// VersionInfo holds version information
type VersionInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"build_time"`
}

// ChatMessage represents a message in the conversation
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest for LLM inference
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

// ChatResponse from throne
type ChatResponse struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Model     string `json:"model"`
	Success   bool   `json:"success"`
	LatencyMs int64  `json:"latency_ms"`
}

// ModelInfo represents an available model
type ModelInfo struct {
	Name string `json:"name"`
	Size string `json:"size,omitempty"`
}

// GetVersion fetches throne daemon version
func (c *ThroneClient) GetVersion() (*VersionInfo, error) {
	resp, err := c.client.Get(c.BaseURL + "/version")
	if err != nil {
		return nil, fmt.Errorf("failed to get version: %w", err)
	}
	defer resp.Body.Close()

	var version VersionInfo
	if err := json.NewDecoder(resp.Body).Decode(&version); err != nil {
		return nil, fmt.Errorf("failed to decode version: %w", err)
	}

	return &version, nil
}

// Chat sends a chat request to throne
func (c *ThroneClient) Chat(model, prompt string) (*ChatResponse, error) {
	reqBody, err := json.Marshal(ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{Role: "user", Content: prompt},
		},
		Stream: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Post(c.BaseURL+"/chat", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to send chat request: %w", err)
	}
	defer resp.Body.Close()

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chatResp, nil
}

// ListModels fetches available Ollama models
func (c *ThroneClient) ListModels() ([]string, error) {
	resp, err := c.client.Get(c.BaseURL + "/ollama/models")
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %w", err)
	}
	defer resp.Body.Close()

	var models []string
	if err := json.NewDecoder(resp.Body).Decode(&models); err != nil {
		return nil, fmt.Errorf("failed to decode models: %w", err)
	}

	return models, nil
}

// Health checks throne daemon health
func (c *ThroneClient) Health() error {
	resp, err := c.client.Get(c.BaseURL + "/healthz")
	if err != nil {
		return fmt.Errorf("throne daemon not responding: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("throne daemon unhealthy: %s", string(body))
	}

	return nil
}
