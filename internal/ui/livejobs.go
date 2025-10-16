package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sovereynai/reign/internal/client"
	"github.com/sovereynai/reign/internal/config"
)

// Job represents a running inference job
type Job struct {
	ID        string
	Model     string
	ModelType string // "ollama" or "onnx"
	Status    string // "queued", "running", "completed", "failed"
	Progress  float64 // 0.0 to 1.0
	StartTime time.Time
	Duration  time.Duration
	NodeID    string
}

type liveJobsModel struct {
	jobs       []Job
	spinner    spinner.Model
	progress   progress.Model
	quitting   bool
	lastUpdate time.Time
}

type jobUpdateMsg struct {
	jobs []Job
}

type jobsTickMsg time.Time

func tickEvery(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(t time.Time) tea.Msg {
		return jobsTickMsg(t)
	})
}

func InitialJobsModel() liveJobsModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := progress.New(progress.WithDefaultGradient())

	return liveJobsModel{
		jobs:       []Job{},
		spinner:    s,
		progress:   p,
		lastUpdate: time.Now(),
	}
}

func (m liveJobsModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tickEvery(500*time.Millisecond),
	)
}

func (m liveJobsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "r":
			// Refresh - fetch new jobs
			return m, fetchJobs
		}

	case jobUpdateMsg:
		m.jobs = msg.jobs
		m.lastUpdate = time.Now()
		return m, nil

	case jobsTickMsg:
		// Update job durations and progress
		for i := range m.jobs {
			if m.jobs[i].Status == "running" {
				m.jobs[i].Duration = time.Since(m.jobs[i].StartTime)
				// Simulate progress for running jobs
				if m.jobs[i].Progress < 0.95 {
					m.jobs[i].Progress += 0.05
				}
			}
		}
		return m, tickEvery(500 * time.Millisecond)

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m liveJobsModel) View() string {
	if m.quitting {
		return ""
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("235")).
		Padding(0, 2)

	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	runningStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("220"))

	completedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("46"))

	failedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("196"))

	queuedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("86")).
		Padding(1, 2).
		Width(80)

	var content strings.Builder

	// Header
	content.WriteString(titleStyle.Render("🔄 LIVE INFERENCE JOBS"))
	content.WriteString("\n\n")

	// Stats
	running := 0
	queued := 0
	completed := 0
	failed := 0
	for _, job := range m.jobs {
		switch job.Status {
		case "running":
			running++
		case "queued":
			queued++
		case "completed":
			completed++
		case "failed":
			failed++
		}
	}

	content.WriteString(headerStyle.Render("📊 Status: "))
	content.WriteString(fmt.Sprintf("%s %d  %s %d  %s %d  %s %d\n\n",
		runningStyle.Render("⚡ Running:"), running,
		queuedStyle.Render("⏳ Queued:"), queued,
		completedStyle.Render("✓ Completed:"), completed,
		failedStyle.Render("✗ Failed:"), failed,
	))

	// Jobs list
	if len(m.jobs) == 0 {
		content.WriteString(labelStyle.Render("  No active jobs. Waiting for inference requests...\n"))
		content.WriteString(labelStyle.Render(fmt.Sprintf("  Last checked: %s\n", m.lastUpdate.Format("15:04:05"))))
	} else {
		for _, job := range m.jobs {
			content.WriteString(renderJob(job, m.progress, m.spinner))
			content.WriteString("\n")
		}
	}

	content.WriteString("\n")
	content.WriteString(labelStyle.Render("Press [r] to refresh  [q] to quit"))

	return boxStyle.Render(content.String())
}

func renderJob(job Job, prog progress.Model, spin spinner.Model) string {
	var statusIcon, statusText string
	var statusStyle lipgloss.Style

	switch job.Status {
	case "running":
		statusIcon = spin.View()
		statusText = "RUNNING"
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	case "queued":
		statusIcon = "⏳"
		statusText = "QUEUED"
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	case "completed":
		statusIcon = "✓"
		statusText = "COMPLETED"
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	case "failed":
		statusIcon = "✗"
		statusText = "FAILED"
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	}

	modelIcon := "🤖"
	if job.ModelType == "onnx" {
		modelIcon = "🔬"
	}

	// Format duration
	duration := job.Duration.Round(time.Millisecond)
	durationStr := duration.String()
	if job.Status == "queued" {
		durationStr = "-"
	}

	jobBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(0, 1).
		Width(76)

	var content strings.Builder
	content.WriteString(fmt.Sprintf("%s %s  %s %s  ",
		statusIcon,
		statusStyle.Bold(true).Render(statusText),
		modelIcon,
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")).Render(job.Model),
	))
	content.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(fmt.Sprintf("(%s)", durationStr)))
	content.WriteString("\n")

	content.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(fmt.Sprintf("Job: %s  Node: %s", job.ID[:8], job.NodeID)))
	content.WriteString("\n")

	// Progress bar for running jobs
	if job.Status == "running" {
		prog.Width = 70
		content.WriteString(prog.ViewAs(job.Progress))
		content.WriteString(fmt.Sprintf(" %.0f%%", job.Progress*100))
	}

	return jobBox.Render(content.String())
}

// fetchJobs fetches real jobs from throne API
func fetchJobs() tea.Msg {
	// Get throne daemon URL
	cfg, err := config.Load()
	if err != nil {
		// Return empty jobs on error
		return jobUpdateMsg{jobs: []Job{}}
	}

	// Create throne client
	throneClient := client.NewThroneClient(cfg.ThroneURL)

	// Fetch live jobs
	response, err := throneClient.GetLiveJobs()
	if err != nil {
		// Return empty jobs on error
		return jobUpdateMsg{jobs: []Job{}}
	}

	// Convert API response to UI jobs
	var jobs []Job

	// Add active jobs
	for _, apiJob := range response.Active {
		startTime, _ := time.Parse(time.RFC3339Nano, apiJob.StartTime)
		jobs = append(jobs, Job{
			ID:        apiJob.ID,
			Model:     apiJob.Model,
			ModelType: apiJob.ModelType,
			Status:    apiJob.Status,
			Progress:  apiJob.Progress,
			StartTime: startTime,
			Duration:  time.Duration(apiJob.Duration) * time.Millisecond,
			NodeID:    apiJob.NodeID,
		})
	}

	// Add recent completed jobs
	for _, apiJob := range response.Recent {
		startTime, _ := time.Parse(time.RFC3339Nano, apiJob.StartTime)
		jobs = append(jobs, Job{
			ID:        apiJob.ID,
			Model:     apiJob.Model,
			ModelType: apiJob.ModelType,
			Status:    apiJob.Status,
			Progress:  apiJob.Progress,
			StartTime: startTime,
			Duration:  time.Duration(apiJob.Duration) * time.Millisecond,
			NodeID:    apiJob.NodeID,
		})
	}

	return jobUpdateMsg{jobs: jobs}
}

// ShowLiveJobs displays the live jobs monitor
func ShowLiveJobs() error {
	p := tea.NewProgram(InitialJobsModel())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running live jobs viewer: %w", err)
	}
	return nil
}
