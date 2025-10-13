package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sovereynai/reign/internal/client"
)

var (
	// Box styles
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Background(lipgloss.Color("235")).
			Padding(0, 1).
			Width(70)

	sectionTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("39")).
				MarginTop(1)

	// Data styles
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Bold(true)

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	warningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("220"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("117"))

	// Table styles
	tableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("86"))

	tableCellStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255"))
)

// RenderDeveloperDashboard renders the AI Developer dashboard
func RenderDeveloperDashboard(stats *client.DashboardStats) string {
	var out strings.Builder

	// Header
	out.WriteString(renderHeader("👑 REIGN - AI Developer Dashboard", stats.Version.Version))
	out.WriteString("\n\n")

	dev := stats.Developer

	// Inference Metrics
	out.WriteString(sectionTitleStyle.Render("🤖 INFERENCE METRICS"))
	out.WriteString("\n")
	out.WriteString(renderModelTable(dev.Models))
	out.WriteString("\n")

	// Credits & Usage
	out.WriteString(sectionTitleStyle.Render("💰 CREDITS & USAGE"))
	out.WriteString("\n")
	out.WriteString(renderCredits(&dev.Credits))
	out.WriteString("\n")

	// Performance
	out.WriteString(sectionTitleStyle.Render("⚡ PERFORMANCE"))
	out.WriteString("\n")
	out.WriteString(renderPerformance(&dev.Performance, &dev.Inference))
	out.WriteString("\n")

	// Smart Insights
	if len(dev.Insights) > 0 {
		out.WriteString(sectionTitleStyle.Render("🎯 SMART INSIGHTS"))
		out.WriteString("\n")
		out.WriteString(renderInsights(dev.Insights))
		out.WriteString("\n")
	}

	// Network Health
	out.WriteString(sectionTitleStyle.Render("📊 NETWORK HEALTH"))
	out.WriteString("\n")
	out.WriteString(renderNetwork(&stats.Network))
	out.WriteString("\n")

	// Quick Actions
	out.WriteString(renderQuickActions([]string{
		"reign dev history     - View request history & replay",
		"reign dev optimize    - Get cost reduction suggestions",
		"reign dev playground  - Interactive model testing",
		"reign dev limits      - Check rate limits & quotas",
	}))

	return borderStyle.Render(out.String())
}

// RenderOperatorDashboard renders the Node Operator dashboard
func RenderOperatorDashboard(stats *client.DashboardStats) string {
	var out strings.Builder

	// Header
	out.WriteString(renderHeader("🏛️  THRONE - Node Operator Dashboard", "Uptime: "+stats.Uptime))
	out.WriteString("\n\n")

	op := stats.Operator

	// Earnings & Contribution
	out.WriteString(sectionTitleStyle.Render("💰 EARNINGS & CONTRIBUTION"))
	out.WriteString("\n")
	out.WriteString(renderEarnings(&op.Earnings))
	out.WriteString("\n")

	// Revenue Breakdown
	out.WriteString(sectionTitleStyle.Render("📈 REVENUE BREAKDOWN"))
	out.WriteString("\n")
	out.WriteString(renderRevenueBreakdown(&op.Earnings.Breakdown, op.Earnings.Today))
	out.WriteString("\n")

	// Workload
	out.WriteString(sectionTitleStyle.Render("🔥 WORKLOAD (Last 24h)"))
	out.WriteString("\n")
	out.WriteString(renderWorkload(&op.Workload))
	out.WriteString("\n")

	// Hardware Utilization
	out.WriteString(sectionTitleStyle.Render("🖥️  HARDWARE UTILIZATION"))
	out.WriteString("\n")
	out.WriteString(renderHardware(&op.Hardware))
	out.WriteString("\n")

	// Models Served
	out.WriteString(sectionTitleStyle.Render("📦 MODELS SERVED"))
	out.WriteString("\n")
	out.WriteString(renderModelsServed(op.ModelsServed))
	out.WriteString("\n")

	// Network Participation
	out.WriteString(sectionTitleStyle.Render("🌍 NETWORK PARTICIPATION"))
	out.WriteString("\n")
	out.WriteString(renderOperatorNetwork(&stats.Network, &op.Reputation))
	out.WriteString("\n")

	// Alerts & Optimization
	if len(op.Alerts) > 0 {
		out.WriteString(sectionTitleStyle.Render("⚠️  ALERTS & OPTIMIZATION"))
		out.WriteString("\n")
		out.WriteString(renderAlerts(op.Alerts))
		out.WriteString("\n")
	}

	// Quick Actions
	out.WriteString(renderQuickActions([]string{
		"reign node earnings   - Detailed revenue breakdown & trends",
		"reign node optimize   - Hardware tuning recommendations",
		"reign node models     - Add/remove models based on demand",
		"reign node peers      - Network connections & health",
		"reign node logs       - Real-time inference log stream",
	}))

	return borderStyle.Render(out.String())
}

// Helper rendering functions

func renderHeader(title, subtitle string) string {
	return headerStyle.Render(fmt.Sprintf("%s     %s", title, mutedStyle.Render(subtitle)))
}

func renderModelTable(models []client.ModelUsage) string {
	if len(models) == 0 {
		return mutedStyle.Render("  No inference activity yet")
	}

	var out strings.Builder

	// Table header
	out.WriteString(fmt.Sprintf("  %-15s %-10s %-12s %-10s %-10s\n",
		tableHeaderStyle.Render("Model"),
		tableHeaderStyle.Render("Today"),
		tableHeaderStyle.Render("7d Avg"),
		tableHeaderStyle.Render("Latency"),
		tableHeaderStyle.Render("Credits"),
	))

	// Table rows
	for _, m := range models {
		out.WriteString(fmt.Sprintf("  %-15s %-10s %-12s %-10s %-10s\n",
			tableCellStyle.Render(m.Name),
			tableCellStyle.Render(fmt.Sprintf("%d req", m.RequestsToday)),
			tableCellStyle.Render(fmt.Sprintf("%.1f/day", m.WeekAvg)),
			tableCellStyle.Render(fmt.Sprintf("%dms", m.AvgLatencyMs)),
			tableCellStyle.Render(fmt.Sprintf("%.1fc", m.CreditsSpent)),
		))
	}

	return out.String()
}

func renderCredits(c *client.CreditStats) string {
	trendIcon := "▼"
	trendColor := successStyle
	if c.TrendPercent > 0 {
		trendIcon = "▲"
		trendColor = errorStyle
	}

	trend := trendColor.Render(fmt.Sprintf("%s %.0f%% vs last week", trendIcon, abs(c.TrendPercent)))

	return fmt.Sprintf("  %s  %s      %s\n  %s     %s\n  %s       %s\n  %s       %s\n",
		labelStyle.Render("Balance:"),
		valueStyle.Render(fmt.Sprintf("%.0f credits", c.Balance)),
		mutedStyle.Render(fmt.Sprintf("(-%.0fc today, +%.0fc earned)", c.TodaySpent, c.TodayEarned)),
		labelStyle.Render("Burn Rate:"),
		valueStyle.Render(fmt.Sprintf("~%.1f credits/day", c.BurnRate)),
		labelStyle.Render("Runway:"),
		valueStyle.Render(fmt.Sprintf("%d days at current usage", c.RunwayDays)),
		labelStyle.Render("Cost Trend:"),
		trend,
	)
}

func renderPerformance(p *client.PerformanceStats, inf *client.InferenceStats) string {
	return fmt.Sprintf("  %s  %s  (p50: %dms, p95: %dms, p99: %dms)\n  %s %s  (%d failures in %d requests)\n  %s    %s\n",
		labelStyle.Render("Avg Latency:"),
		valueStyle.Render(fmt.Sprintf("%dms", p.AvgLatencyMs)),
		p.P50LatencyMs, p.P95LatencyMs, p.P99LatencyMs,
		labelStyle.Render("Success Rate:"),
		successStyle.Render(fmt.Sprintf("%.1f%%", inf.SuccessRate)),
		inf.Failures, inf.Total,
		labelStyle.Render("Preferred:"),
		valueStyle.Render(fmt.Sprintf("Local: %.0f%% | Network: %.0f%%", p.LocalPercent, p.NetworkPercent)),
	)
}

func renderInsights(insights []string) string {
	var out strings.Builder
	for _, insight := range insights {
		out.WriteString("  ")
		out.WriteString(infoStyle.Render("→ "))
		out.WriteString(insight)
		out.WriteString("\n")
	}
	return out.String()
}

func renderNetwork(n *client.NetworkStats) string {
	queueBar := renderProgressBar(float64(n.QueueDepth)/50.0, 10)
	return fmt.Sprintf("  %s %s | Models: %d\n  %s     %s | Est Wait: %.1fs\n",
		labelStyle.Render("Available Peers:"),
		valueStyle.Render(fmt.Sprintf("%d online", n.PeersConnected)),
		n.ModelsAvailable,
		labelStyle.Render("Queue Depth:"),
		queueBar,
		n.EstWaitSec,
	)
}

func renderEarnings(e *client.EarningsStats) string {
	trendIcon := "▲"
	trendColor := successStyle
	if e.WeekTrend < 0 {
		trendIcon = "▼"
		trendColor = errorStyle
	}

	return fmt.Sprintf("  %s        %s    %s\n  %s    %s    %s\n  %s     %s    %s\n  %s      %s      %s\n  %s         %s\n",
		labelStyle.Render("Today:"),
		valueStyle.Render(fmt.Sprintf("+%.1f credits", e.Today)),
		mutedStyle.Render(fmt.Sprintf("(Est: $%.2f USD)", e.TodayUSD)),
		labelStyle.Render("This Week:"),
		valueStyle.Render(fmt.Sprintf("+%.1f credits", e.ThisWeek)),
		trendColor.Render(fmt.Sprintf("%s%.0f%% vs last week", trendIcon, abs(e.WeekTrend))),
		labelStyle.Render("All Time:"),
		valueStyle.Render(fmt.Sprintf("%.0f credits", e.AllTime)),
		mutedStyle.Render("(since Jan 3, 2025)"),
		labelStyle.Render("Pending:"),
		valueStyle.Render(fmt.Sprintf("%.1f credits", e.Pending)),
		mutedStyle.Render("(settlement in ~2h)"),
		labelStyle.Render("Rank:"),
		valueStyle.Render(fmt.Sprintf("#%d / %d nodes globally", e.Rank, e.TotalNodes)),
	)
}

func renderRevenueBreakdown(b *client.EarningsBreakdown, total float64) string {
	inferencePercent := (b.Inference / total) * 100
	bandwidthPercent := (b.Bandwidth / total) * 100
	storagePercent := (b.Storage / total) * 100
	bonusPercent := (b.Bonus / total) * 100

	return fmt.Sprintf("  %s    %.1f credits   (%.0f%% of earnings)\n  %s    %.1f credits   (%.0f%% of earnings)\n  %s      %.1f credits   (%.0f%% of earnings)\n  %s        %.1f credits   (%.0f%% of earnings)\n",
		labelStyle.Render("Inference:"),
		b.Inference, inferencePercent,
		labelStyle.Render("Bandwidth:"),
		b.Bandwidth, bandwidthPercent,
		labelStyle.Render("Storage:"),
		b.Storage, storagePercent,
		labelStyle.Render("Bonus:"),
		b.Bonus, bonusPercent,
	)
}

func renderWorkload(w *client.WorkloadStats) string {
	return fmt.Sprintf("  %s  %s\n  %s     %s  (%d failures)\n  %s      %s (target: <150ms for premium tier)\n  %s        %s - %d req/hr\n",
		labelStyle.Render("Requests Served:"),
		valueStyle.Render(fmt.Sprintf("%d inferences", w.RequestsServed)),
		labelStyle.Render("Success Rate:"),
		successStyle.Render(fmt.Sprintf("%.1f%%", w.SuccessRate)),
		w.Failures,
		labelStyle.Render("Avg Latency:"),
		valueStyle.Render(fmt.Sprintf("%dms", w.AvgLatencyMs)),
		labelStyle.Render("Peak Hour:"),
		valueStyle.Render(w.PeakHour),
		w.PeakRequests,
	)
}

func renderHardware(h *client.HardwareStats) string {
	gpuBar := renderProgressBar(h.GPU.Percent/100.0, 10)
	cpuBar := renderProgressBar(h.CPU.Percent/100.0, 10)
	ramBar := renderProgressBar(h.RAM.Percent/100.0, 10)
	diskBar := renderProgressBar(h.Disk.Percent/100.0, 10)

	tempColor := successStyle
	if h.Temperature.GPU > 75 {
		tempColor = warningStyle
	}
	if h.Temperature.GPU > 85 {
		tempColor = errorStyle
	}

	return fmt.Sprintf("  %s   %s %.0f%% (%s)\n  %s   %s %.0f%% (%s)\n  %s   %s %.0f%% (%s)\n  %s  %s %.0f%% (%s)\n  %s  GPU: %s  |  CPU: %s  %s\n  %s ~%.0fW avg  (est $%.2f/day @ $0.12/kWh)\n",
		labelStyle.Render("GPU:"),
		gpuBar, h.GPU.Percent, mutedStyle.Render(h.GPU.Details),
		labelStyle.Render("CPU:"),
		cpuBar, h.CPU.Percent, mutedStyle.Render(h.CPU.Details),
		labelStyle.Render("RAM:"),
		ramBar, h.RAM.Percent, mutedStyle.Render(fmt.Sprintf("%s / %s", h.RAM.Used, h.RAM.Total)),
		labelStyle.Render("Disk:"),
		diskBar, h.Disk.Percent, mutedStyle.Render(fmt.Sprintf("%s / %s", h.Disk.Used, h.Disk.Total)),
		labelStyle.Render("Temp:"),
		tempColor.Render(fmt.Sprintf("%.0f°C", h.Temperature.GPU)),
		tempColor.Render(fmt.Sprintf("%.0f°C", h.Temperature.CPU)),
		mutedStyle.Render(fmt.Sprintf("(%s)", h.Temperature.Status)),
		labelStyle.Render("Power:"),
		h.PowerWatts, h.PowerCostDaily,
	)
}

func renderModelsServed(models []client.ModelServed) string {
	if len(models) == 0 {
		return mutedStyle.Render("  No models being served")
	}

	var out strings.Builder
	for _, m := range models {
		statusIcon := "✓"
		statusColor := successStyle
		if m.Status == "warning" {
			statusIcon = "⚠"
			statusColor = warningStyle
		} else if m.Status == "idle" {
			statusIcon = "○"
			statusColor = mutedStyle
		}

		out.WriteString(fmt.Sprintf("  %s %-15s %d reqs  |  Avg: %dms   |  Rev: %.1fc\n",
			statusColor.Render(statusIcon),
			m.Name,
			m.Requests,
			m.AvgLatencyMs,
			m.Revenue,
		))
	}
	return out.String()
}

func renderOperatorNetwork(n *client.NetworkStats, r *client.ReputationStats) string {
	stars := renderStars(r.Score, r.MaxScore)
	return fmt.Sprintf("  %s  %s\n  %s     %.2f GB (earning +%.1fc)\n  %s      %d in network | 3 assigned to you\n  %s       %s %.1f/%.0f (%d ratings)\n",
		labelStyle.Render("Peers Connected:"),
		valueStyle.Render(fmt.Sprintf("%d nodes", n.PeersConnected)),
		labelStyle.Render("Data Relayed:"),
		n.DataRelayedGB, n.DataRelayedGB*3.5,
		labelStyle.Render("Jobs Queued:"),
		n.QueueDepth,
		labelStyle.Render("Reputation:"),
		stars,
		r.Score, r.MaxScore, r.Ratings,
	)
}

func renderAlerts(alerts []client.Alert) string {
	var out strings.Builder
	for _, alert := range alerts {
		icon := "→"
		style := infoStyle
		if alert.Level == "warning" {
			icon = "→"
			style = warningStyle
		} else if alert.Level == "error" {
			icon = "✗"
			style = errorStyle
		} else if alert.Level == "info" {
			icon = "✓"
			style = successStyle
		}

		out.WriteString("  ")
		out.WriteString(style.Render(icon + " "))
		out.WriteString(alert.Message)
		out.WriteString("\n")
	}
	return out.String()
}

func renderQuickActions(actions []string) string {
	var out strings.Builder
	out.WriteString("\n")
	out.WriteString(mutedStyle.Render("Quick Actions:"))
	out.WriteString("\n")
	for _, action := range actions {
		out.WriteString("  ")
		out.WriteString(mutedStyle.Render(action))
		out.WriteString("\n")
	}
	return out.String()
}

func renderProgressBar(percent float64, width int) string {
	if percent > 1.0 {
		percent = 1.0
	}
	if percent < 0 {
		percent = 0
	}

	filled := int(percent * float64(width))
	empty := width - filled

	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)

	var color lipgloss.Style
	if percent < 0.5 {
		color = successStyle
	} else if percent < 0.8 {
		color = warningStyle
	} else {
		color = errorStyle
	}

	return color.Render("[" + bar + "]")
}

func renderStars(score, maxScore float64) string {
	rating := (score / maxScore) * 5.0
	fullStars := int(rating)
	halfStar := rating-float64(fullStars) >= 0.5
	emptyStars := 5 - fullStars
	if halfStar {
		emptyStars--
	}

	stars := strings.Repeat("★", fullStars)
	if halfStar {
		stars += "☆"
	}
	stars += strings.Repeat("☆", emptyStars)

	return warningStyle.Render(stars)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
