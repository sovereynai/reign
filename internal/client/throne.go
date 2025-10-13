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

// DashboardStats contains comprehensive stats for dashboards
type DashboardStats struct {
	Role            string              `json:"role"` // "developer", "operator", or "both"
	Uptime          string              `json:"uptime"`
	Developer       *DeveloperStats     `json:"developer,omitempty"`
	Operator        *OperatorStats      `json:"operator,omitempty"`
	Network         NetworkStats        `json:"network"`
	Version         VersionInfo         `json:"version"`
}

// DeveloperStats for AI developers
type DeveloperStats struct {
	Credits         CreditStats         `json:"credits"`
	Inference       InferenceStats      `json:"inference"`
	Performance     PerformanceStats    `json:"performance"`
	Models          []ModelUsage        `json:"models"`
	Insights        []string            `json:"insights"`
}

// OperatorStats for node operators
type OperatorStats struct {
	Earnings        EarningsStats       `json:"earnings"`
	Workload        WorkloadStats       `json:"workload"`
	Hardware        HardwareStats       `json:"hardware"`
	ModelsServed    []ModelServed       `json:"models_served"`
	Reputation      ReputationStats     `json:"reputation"`
	Alerts          []Alert             `json:"alerts"`
}

// CreditStats tracks credit usage
type CreditStats struct {
	Balance         float64 `json:"balance"`
	TodaySpent      float64 `json:"today_spent"`
	TodayEarned     float64 `json:"today_earned"`
	BurnRate        float64 `json:"burn_rate"`
	RunwayDays      int     `json:"runway_days"`
	TrendPercent    float64 `json:"trend_percent"`
}

// InferenceStats tracks inference usage
type InferenceStats struct {
	Today           int     `json:"today"`
	WeekAvg         float64 `json:"week_avg"`
	Total           int     `json:"total"`
	SuccessRate     float64 `json:"success_rate"`
	Failures        int     `json:"failures"`
}

// PerformanceStats tracks latency
type PerformanceStats struct {
	AvgLatencyMs    int     `json:"avg_latency_ms"`
	P50LatencyMs    int     `json:"p50_latency_ms"`
	P95LatencyMs    int     `json:"p95_latency_ms"`
	P99LatencyMs    int     `json:"p99_latency_ms"`
	LocalPercent    float64 `json:"local_percent"`
	NetworkPercent  float64 `json:"network_percent"`
}

// ModelUsage tracks per-model usage
type ModelUsage struct {
	Name            string  `json:"name"`
	RequestsToday   int     `json:"requests_today"`
	WeekAvg         float64 `json:"week_avg"`
	AvgLatencyMs    int     `json:"avg_latency_ms"`
	CreditsSpent    float64 `json:"credits_spent"`
}

// EarningsStats tracks operator earnings
type EarningsStats struct {
	Today           float64 `json:"today"`
	TodayUSD        float64 `json:"today_usd"`
	ThisWeek        float64 `json:"this_week"`
	WeekTrend       float64 `json:"week_trend"`
	AllTime         float64 `json:"all_time"`
	Pending         float64 `json:"pending"`
	Held            float64 `json:"held"` // Amount held for security/escrow
	Rank            int     `json:"rank"`
	TotalNodes      int     `json:"total_nodes"`
	Breakdown       EarningsBreakdown `json:"breakdown"`
}

// EarningsBreakdown shows revenue sources
type EarningsBreakdown struct {
	Inference       float64 `json:"inference"`
	Bandwidth       float64 `json:"bandwidth"`
	Storage         float64 `json:"storage"`
	Bonus           float64 `json:"bonus"`
}

// WorkloadStats tracks node workload
type WorkloadStats struct {
	RequestsServed  int     `json:"requests_served"`
	SuccessRate     float64 `json:"success_rate"`
	Failures        int     `json:"failures"`
	AvgLatencyMs    int     `json:"avg_latency_ms"`
	PeakHour        string  `json:"peak_hour"`
	PeakRequests    int     `json:"peak_requests"`
}

// HardwareStats tracks hardware utilization
type HardwareStats struct {
	GPU             ResourceUsage   `json:"gpu"`
	CPU             ResourceUsage   `json:"cpu"`
	RAM             ResourceUsage   `json:"ram"`
	Disk            ResourceUsage   `json:"disk"`
	Temperature     TempStats       `json:"temperature"`
	PowerWatts      float64         `json:"power_watts"`
	PowerCostDaily  float64         `json:"power_cost_daily"`
}

// ResourceUsage for hardware resources
type ResourceUsage struct {
	Percent         float64 `json:"percent"`
	Used            string  `json:"used"`
	Total           string  `json:"total"`
	Details         string  `json:"details"`
}

// TempStats for temperature monitoring
type TempStats struct {
	GPU             float64 `json:"gpu"`
	CPU             float64 `json:"cpu"`
	Status          string  `json:"status"` // "healthy", "warm", "hot"
}

// ModelServed tracks models being served
type ModelServed struct {
	Name            string  `json:"name"`
	Requests        int     `json:"requests"`
	AvgLatencyMs    int     `json:"avg_latency_ms"`
	Revenue         float64 `json:"revenue"`
	Status          string  `json:"status"` // "active", "idle", "warning"
}

// ReputationStats tracks node reputation
type ReputationStats struct {
	Score           float64 `json:"score"`
	MaxScore        float64 `json:"max_score"`
	Ratings         int     `json:"ratings"`
	UptimeStreak    int     `json:"uptime_streak_days"`
}

// Alert represents system alerts
type Alert struct {
	Level           string  `json:"level"` // "info", "warning", "error"
	Message         string  `json:"message"`
}

// NetworkStats tracks network health
type NetworkStats struct {
	PeersConnected  int     `json:"peers_connected"`
	ModelsAvailable int     `json:"models_available"`
	QueueDepth      int     `json:"queue_depth"`
	EstWaitSec      float64 `json:"est_wait_sec"`
	DataRelayedGB   float64 `json:"data_relayed_gb"`
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

// GetDashboardStats fetches comprehensive dashboard statistics
func (c *ThroneClient) GetDashboardStats() (*DashboardStats, error) {
	resp, err := c.client.Get(c.BaseURL + "/stats/dashboard")
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard stats: %w", err)
	}
	defer resp.Body.Close()

	var stats DashboardStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to decode dashboard stats: %w", err)
	}

	return &stats, nil
}
