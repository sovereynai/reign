package config

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Config holds reign CLI configuration
type Config struct {
	ThroneURL string
}

// Load discovers the throne daemon URL
func Load() (*Config, error) {
	// Check environment variable first
	if url := os.Getenv("THRONE_URL"); url != "" {
		return &Config{ThroneURL: url}, nil
	}

	// Try localhost:8080 (default throne port)
	defaultURLs := []string{
		"http://localhost:8080",
		"http://localhost:8090",
		"http://localhost:8091",
	}

	for _, url := range defaultURLs {
		if isReachable(url) {
			return &Config{ThroneURL: url}, nil
		}
	}

	return nil, fmt.Errorf("throne daemon not found. Start with: throne serve")
}

func isReachable(url string) bool {
	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url + "/healthz")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
