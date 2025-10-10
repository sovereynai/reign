package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ChatMessage represents a single message in a chat conversation
type ChatMessage struct {
	Role    string `json:"role"`    // "user" or "assistant"
	Content string `json:"content"`
}

// ChatRequest represents a chat request
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

// ChatResponse represents a chat response
type ChatResponse struct {
	Message         ChatMessage `json:"message"`
	Model           string      `json:"model"`
	Success         bool        `json:"success"`
	TokensGenerated int         `json:"tokens_generated"`
	LatencyMs       int64       `json:"latency_ms"`
}

// Chat sends a conversational request to the daemon
func Chat(model string, messages []ChatMessage) (*ChatResponse, error) {
	req := ChatRequest{
		Model:    model,
		Messages: messages,
		Stream:   false,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	resp, err := http.Post(BaseURL+"/chat", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("post chat: %w", err)
	}
	defer resp.Body.Close()

	rspBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("chat failed: %s", rspBody)
	}

	var out ChatResponse
	if err := json.Unmarshal(rspBody, &out); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &out, nil
}
