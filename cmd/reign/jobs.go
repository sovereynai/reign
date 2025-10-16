package main

import (
	"github.com/sovereynai/reign/internal/ui"
	"github.com/spf13/cobra"
)

var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "View live inference jobs with real-time progress",
	Long: `Display a real-time dashboard of inference jobs being processed by your throne node.
Shows running, queued, completed, and failed jobs with progress bars and timing information.

This is particularly useful for node operators to monitor their workload in real-time.

Example:
  reign jobs           # Start live jobs monitor
  reign node jobs      # Same thing (alias)
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.ShowLiveJobs(); err != nil {
			// If Bubble Tea fails (no TTY), show a message
			println("❌ Live jobs viewer requires a terminal (TTY)")
			println("💡 Tip: Use 'reign node status' for a snapshot view")
		}
	},
}

func init() {
	jobsCmd.Flags().BoolP("watch", "w", false, "Watch mode (continuous updates)")
	jobsCmd.Flags().IntP("refresh", "n", 1, "Refresh interval in seconds")
}

func RegisterJobsCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(jobsCmd)
}
