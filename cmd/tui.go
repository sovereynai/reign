package cmd

import (
	"log"

	"github.com/Leathal1/greycli/internal/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the interactive terminal UI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.Run(); err != nil {
			log.Fatalf("Failed to run TUI: %v", err)
		}
	},
}
