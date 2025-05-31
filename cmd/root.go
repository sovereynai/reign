package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "greycli",
	Short: "greycli is the terminal UI and CLI for interacting with the Fist Network",
	Long:  `greycli provides a terminal interface and CLI tools to submit jobs, view models, credits, and interact with the Fist Network daemon.`,
}

// Execute runs the root command
func Execute() {
	rootCmd.AddCommand(tuiCmd) // Add subcommands like 'tui' here
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
