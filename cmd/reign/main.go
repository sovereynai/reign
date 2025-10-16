package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sovereynai/reign/internal/bootstrap"
	"github.com/sovereynai/reign/internal/client"
	"github.com/sovereynai/reign/internal/config"
	"github.com/sovereynai/reign/internal/ui"
	"github.com/spf13/cobra"
)

var (
	// Styles
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginTop(1).
			MarginBottom(1)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))
)

func main() {
	// First-run setup (only happens once)
	if err := bootstrap.Setup(); err != nil {
		// Setup failed, but don't block - user might be fixing issues manually
		// Error already displayed by Setup()
	}

	rootCmd := &cobra.Command{
		Use:   "reign",
		Short: "Sovereyn CLI - Interface for distributed AI inference",
		Long: titleStyle.Render("👑 Reign") + "\n\n" +
			"The command-line interface for Sovereyn's distributed intelligence network.\n" +
			"Submit inference jobs, manage models, and monitor the network.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Skip throne check for help/version commands
			if cmd.Name() == "version" || cmd.Name() == "help" || cmd.Name() == "completion" {
				return nil
			}
			// Ensure throne daemon is running
			return bootstrap.EnsureThroneRunning()
		},
	}

	// Version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		RunE:  runVersion,
	}

	// Chat command
	chatCmd := &cobra.Command{
		Use:   "chat [prompt]",
		Short: "Chat with an AI model",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runChat,
	}
	chatCmd.Flags().StringP("model", "m", "llama3.2:3b", "Model to use for inference")

	// Models command
	modelsCmd := &cobra.Command{
		Use:   "models",
		Short: "List available models",
		RunE:  runModels,
	}
	modelsNetworkCmd := &cobra.Command{
		Use:   "network",
		Short: "List all models available across the network",
		RunE:  runModelsNetwork,
	}
	modelsLocateCmd := &cobra.Command{
		Use:   "locate [model]",
		Short: "Find where a specific model is available",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runModelsLocate,
	}
	modelsCmd.AddCommand(modelsNetworkCmd, modelsLocateCmd)

	// Status command (enhanced)
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show comprehensive dashboard (auto-detects role)",
		RunE:  runStatus,
	}

	// Dev subcommand
	devCmd := &cobra.Command{
		Use:   "dev",
		Short: "AI Developer commands and dashboard",
	}
	devStatusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show AI Developer dashboard",
		RunE:  runDevStatus,
	}
	devHistoryCmd := &cobra.Command{
		Use:   "history",
		Short: "View request history (coming soon)",
		RunE:  runComingSoon,
	}
	devOptimizeCmd := &cobra.Command{
		Use:   "optimize",
		Short: "Get cost optimization suggestions (coming soon)",
		RunE:  runComingSoon,
	}
	devPlaygroundCmd := &cobra.Command{
		Use:   "playground",
		Short: "Interactive model testing (coming soon)",
		RunE:  runComingSoon,
	}
	devCmd.AddCommand(devStatusCmd, devHistoryCmd, devOptimizeCmd, devPlaygroundCmd)

	// Node subcommand
	nodeCmd := &cobra.Command{
		Use:   "node",
		Short: "Node Operator commands and dashboard",
	}
	nodeStatusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show Node Operator dashboard",
		RunE:  runNodeStatus,
	}
	nodeEarningsCmd := &cobra.Command{
		Use:   "earnings",
		Short: "Detailed revenue breakdown (coming soon)",
		RunE:  runComingSoon,
	}
	nodeOptimizeCmd := &cobra.Command{
		Use:   "optimize",
		Short: "Hardware optimization suggestions (coming soon)",
		RunE:  runComingSoon,
	}
	nodeModelsCmd := &cobra.Command{
		Use:   "models",
		Short: "Manage models based on demand (coming soon)",
		RunE:  runComingSoon,
	}
	nodePeersCmd := &cobra.Command{
		Use:   "peers",
		Short: "Network connections & health (coming soon)",
		RunE:  runComingSoon,
	}
	nodeLogsCmd := &cobra.Command{
		Use:   "logs",
		Short: "Real-time inference logs (coming soon)",
		RunE:  runComingSoon,
	}
	nodeJobsCmd := &cobra.Command{
		Use:   "jobs",
		Short: "View live inference jobs with progress bars",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ui.ShowLiveJobs(); err != nil {
				fmt.Println(errorStyle.Render("❌ Live jobs viewer requires a terminal (TTY)"))
				fmt.Println(infoStyle.Render("💡 Tip: Use 'reign node status' for a snapshot view"))
			}
		},
	}
	nodeCmd.AddCommand(nodeStatusCmd, nodeEarningsCmd, nodeOptimizeCmd, nodeModelsCmd, nodePeersCmd, nodeLogsCmd, nodeJobsCmd)

	// Register jobs command (also available as top-level command)
	RegisterJobsCommand(rootCmd)

	rootCmd.AddCommand(versionCmd, chatCmd, modelsCmd, statusCmd, devCmd, nodeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, errorStyle.Render("Error: "+err.Error()))
		os.Exit(1)
	}
}

func getThroneClient() (*client.ThroneClient, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return client.NewThroneClient(cfg.ThroneURL), nil
}

func runVersion(cmd *cobra.Command, args []string) error {
	c, err := getThroneClient()
	if err != nil {
		return err
	}

	version, err := c.GetVersion()
	if err != nil {
		return fmt.Errorf("failed to get version: %w", err)
	}

	fmt.Println(titleStyle.Render("👑 Sovereyn"))
	fmt.Println(infoStyle.Render("Daemon Version:  ") + version.Version)
	fmt.Println(infoStyle.Render("Commit:          ") + version.Commit[:8])
	fmt.Println(infoStyle.Render("Build Time:      ") + version.BuildTime)
	fmt.Println(infoStyle.Render("CLI Version:     ") + "v0.2.1")

	return nil
}

func runChat(cmd *cobra.Command, args []string) error {
	model, _ := cmd.Flags().GetString("model")
	prompt := strings.Join(args, " ")

	c, err := getThroneClient()
	if err != nil {
		return err
	}

	// Show we're working
	fmt.Println(infoStyle.Render("🤖 Submitting to throne daemon..."))
	fmt.Println(infoStyle.Render("📝 Model: ") + model)
	fmt.Println(infoStyle.Render("💬 Prompt: ") + prompt)
	fmt.Println()

	resp, err := c.Chat(model, prompt)
	if err != nil {
		return fmt.Errorf("inference failed: %w", err)
	}

	if !resp.Success {
		return fmt.Errorf("inference returned success=false")
	}

	// Display response
	fmt.Println(titleStyle.Render("✨ Response"))
	if resp.Message.Content != "" {
		fmt.Println(resp.Message.Content)
	} else {
		fmt.Println(infoStyle.Render("(Inference completed but response was empty)"))
	}

	fmt.Println()
	fmt.Println(infoStyle.Render(fmt.Sprintf("⚡ Latency: %dms", resp.LatencyMs)))

	return nil
}

func runModels(cmd *cobra.Command, args []string) error {
	c, err := getThroneClient()
	if err != nil {
		return err
	}

	models, err := c.ListModels()
	if err != nil {
		return fmt.Errorf("failed to list models: %w", err)
	}

	fmt.Println(titleStyle.Render("📦 Local Models"))
	for _, model := range models {
		fmt.Println(successStyle.Render("  • ") + model)
	}

	return nil
}

func runModelsNetwork(cmd *cobra.Command, args []string) error {
	c, err := getThroneClient()
	if err != nil {
		return err
	}

	models, err := c.ListNetworkModels()
	if err != nil {
		return fmt.Errorf("failed to list network models: %w", err)
	}

	fmt.Println(titleStyle.Render("🌍 Network Models"))
	fmt.Println()

	if len(models) == 0 {
		fmt.Println(infoStyle.Render("No models found on the network"))
		return nil
	}

	// Group by type
	ollamaModels := []client.NetworkModel{}
	onnxModels := []client.NetworkModel{}
	for _, model := range models {
		if model.Type == client.ModelTypeOllama {
			ollamaModels = append(ollamaModels, model)
		} else {
			onnxModels = append(onnxModels, model)
		}
	}

	// Display Ollama models
	if len(ollamaModels) > 0 {
		fmt.Println(successStyle.Render(fmt.Sprintf("🤖 Ollama Models (%d)", len(ollamaModels))))
		for _, model := range ollamaModels {
			fmt.Printf("  %s %s\n",
				successStyle.Render("• "+model.Name),
				infoStyle.Render(fmt.Sprintf("(%s)", model.Category)))

			trackerInfo := ""
			if model.TrackerCount > 0 {
				trackerInfo = fmt.Sprintf("%d tracker(s), ", model.TrackerCount)
			}
			fmt.Println(infoStyle.Render(fmt.Sprintf("    %s%d location(s)", trackerInfo, model.PeerCount)))

			// Show tracker locations
			trackers := []client.ModelLocation{}
			remotes := []client.ModelLocation{}
			var localLoc *client.ModelLocation

			for _, loc := range model.Locations {
				if loc.Type == "tracker" {
					trackers = append(trackers, loc)
				} else if loc.Type == "local" {
					localLoc = &loc
				} else {
					remotes = append(remotes, loc)
				}
			}

			if localLoc != nil {
				fmt.Println(infoStyle.Render("      ✓ Local (this node)"))
			}
			for _, tracker := range trackers {
				fmt.Println(infoStyle.Render(fmt.Sprintf("      📍 Tracker: %s", tracker.TrackerName)))
			}
			for _, remote := range remotes {
				fmt.Println(infoStyle.Render(fmt.Sprintf("      🌐 Node: %s", remote.NodeID)))
			}

			if model.PullCommand != "" {
				fmt.Println(infoStyle.Render(fmt.Sprintf("    💾 Pull: %s", model.PullCommand)))
			}
			fmt.Println()
		}
	}

	// Display ONNX models
	if len(onnxModels) > 0 {
		fmt.Println(successStyle.Render(fmt.Sprintf("🔬 ONNX Models (%d)", len(onnxModels))))
		for _, model := range onnxModels {
			fmt.Printf("  %s %s\n",
				successStyle.Render("• "+model.Name),
				infoStyle.Render(fmt.Sprintf("(%s)", model.Category)))

			trackerInfo := ""
			if model.TrackerCount > 0 {
				trackerInfo = fmt.Sprintf("%d tracker(s), ", model.TrackerCount)
			}
			fmt.Println(infoStyle.Render(fmt.Sprintf("    %s%d location(s)", trackerInfo, model.PeerCount)))

			// Show tracker locations
			trackers := []client.ModelLocation{}
			remotes := []client.ModelLocation{}
			var localLoc *client.ModelLocation

			for _, loc := range model.Locations {
				if loc.Type == "tracker" {
					trackers = append(trackers, loc)
				} else if loc.Type == "local" {
					localLoc = &loc
				} else {
					remotes = append(remotes, loc)
				}
			}

			if localLoc != nil {
				fmt.Println(infoStyle.Render("      ✓ Local (this node)"))
			}
			for _, tracker := range trackers {
				fmt.Println(infoStyle.Render(fmt.Sprintf("      📍 Tracker: %s", tracker.TrackerName)))
			}
			for _, remote := range remotes {
				fmt.Println(infoStyle.Render(fmt.Sprintf("      🌐 Node: %s", remote.NodeID)))
			}

			if model.PullCommand != "" {
				fmt.Println(infoStyle.Render(fmt.Sprintf("    💾 Pull: %s", model.PullCommand)))
			}
			fmt.Println()
		}
	}

	fmt.Println(infoStyle.Render(fmt.Sprintf("Total: %d models (%d Ollama, %d ONNX)",
		len(models), len(ollamaModels), len(onnxModels))))

	return nil
}

func runModelsLocate(cmd *cobra.Command, args []string) error {
	modelName := args[0]

	c, err := getThroneClient()
	if err != nil {
		return err
	}

	locations, err := c.LocateModel(modelName)
	if err != nil {
		return fmt.Errorf("failed to locate model: %w", err)
	}

	fmt.Println(titleStyle.Render(fmt.Sprintf("📍 Locations for %s", modelName)))
	fmt.Println()

	if len(locations) == 0 {
		fmt.Println(errorStyle.Render("Model not found on the network"))
		return nil
	}

	for _, loc := range locations {
		nodeID := loc["node_id"].(string)
		locType := loc["type"].(string)

		if locType == "local" {
			fmt.Println(successStyle.Render("  ✓ Local (this node)"))
			fmt.Println(infoStyle.Render(fmt.Sprintf("    URL: %s", loc["url"])))
		} else {
			fmt.Println(successStyle.Render(fmt.Sprintf("  ✓ Remote node: %s", nodeID)))
			fmt.Println(infoStyle.Render(fmt.Sprintf("    URL: %s", loc["url"])))
			if latency, ok := loc["latency_ms"]; ok && latency.(float64) > 0 {
				fmt.Println(infoStyle.Render(fmt.Sprintf("    Latency: %.0fms", latency)))
			}
		}
	}

	return nil
}

func runStatus(cmd *cobra.Command, args []string) error {
	c, err := getThroneClient()
	if err != nil {
		return err
	}

	stats, err := c.GetDashboardStats()
	if err != nil {
		// Fallback to simple status if dashboard endpoint not available
		return runSimpleStatus(c)
	}

	// Auto-detect role and show appropriate dashboard
	switch stats.Role {
	case "developer":
		fmt.Println(ui.RenderDeveloperDashboard(stats))
	case "operator":
		fmt.Println(ui.RenderOperatorDashboard(stats))
	case "both":
		// Show both dashboards
		fmt.Println(ui.RenderDeveloperDashboard(stats))
		fmt.Println()
		fmt.Println(ui.RenderOperatorDashboard(stats))
	default:
		return runSimpleStatus(c)
	}

	return nil
}

func runDevStatus(cmd *cobra.Command, args []string) error {
	c, err := getThroneClient()
	if err != nil {
		return err
	}

	stats, err := c.GetDashboardStats()
	if err != nil {
		return fmt.Errorf("failed to get dashboard stats: %w", err)
	}

	if stats.Developer == nil {
		return fmt.Errorf("no developer stats available - have you made any inference requests?")
	}

	fmt.Println(ui.RenderDeveloperDashboard(stats))
	return nil
}

func runNodeStatus(cmd *cobra.Command, args []string) error {
	c, err := getThroneClient()
	if err != nil {
		return err
	}

	stats, err := c.GetDashboardStats()
	if err != nil {
		return fmt.Errorf("failed to get dashboard stats: %w", err)
	}

	if stats.Operator == nil {
		return fmt.Errorf("no operator stats available - is this node serving models?")
	}

	fmt.Println(ui.RenderOperatorDashboard(stats))
	return nil
}

func runComingSoon(cmd *cobra.Command, args []string) error {
	fmt.Println(infoStyle.Render("🚧 Coming soon!"))
	fmt.Println("This feature is under active development.")
	return nil
}

// Fallback for older throne versions without dashboard endpoint
func runSimpleStatus(c *client.ThroneClient) error {
	if err := c.Health(); err != nil {
		fmt.Println(errorStyle.Render("❌ Throne daemon: OFFLINE"))
		return err
	}

	version, err := c.GetVersion()
	if err != nil {
		return err
	}

	fmt.Println(titleStyle.Render("🏛️  Throne Daemon Status"))
	fmt.Println(successStyle.Render("✅ Status:  ") + "ONLINE")
	fmt.Println(infoStyle.Render("🔖 Version: ") + version.Version)
	fmt.Println(infoStyle.Render("🌐 URL:     ") + c.BaseURL)

	// Show available models
	models, err := c.ListModels()
	if err == nil && len(models) > 0 {
		fmt.Println()
		fmt.Println(infoStyle.Render("📦 Models:  ") + fmt.Sprintf("%d available", len(models)))
	}

	fmt.Println()
	fmt.Println(infoStyle.Render("💡 Tip: Upgrade throne for rich dashboards with detailed metrics"))

	return nil
}
