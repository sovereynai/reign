package main

import (
	"fmt"

	"github.com/sovereynai/reign/internal/client"
	"github.com/sovereynai/reign/internal/ui"
	"github.com/spf13/cobra"
)

func createDemoCommand() *cobra.Command {
	demoCmd := &cobra.Command{
		Use:    "demo",
		Short:  "Show demo dashboards with mock data",
		Hidden: true, // Hidden from help - just for testing/demos
	}

	demoDev := &cobra.Command{
		Use:   "dev",
		Short: "Show AI Developer dashboard with mock data",
		RunE: func(cmd *cobra.Command, args []string) error {
			stats := client.MockDeveloperStats()
			fmt.Println(ui.RenderDeveloperDashboard(stats))
			return nil
		},
	}

	demoNode := &cobra.Command{
		Use:   "node",
		Short: "Show Node Operator dashboard with mock data",
		RunE: func(cmd *cobra.Command, args []string) error {
			stats := client.MockOperatorStats()
			fmt.Println(ui.RenderOperatorDashboard(stats))
			return nil
		},
	}

	demoBoth := &cobra.Command{
		Use:   "both",
		Short: "Show both dashboards with mock data",
		RunE: func(cmd *cobra.Command, args []string) error {
			stats := client.MockBothStats()
			fmt.Println(ui.RenderDeveloperDashboard(stats))
			fmt.Println()
			fmt.Println(ui.RenderOperatorDashboard(stats))
			return nil
		},
	}

	demoCmd.AddCommand(demoDev, demoNode, demoBoth)
	return demoCmd
}
