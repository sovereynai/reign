package cmd

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Knuckle daemon",
	Long:  "Stops the running Knuckle daemon gracefully.",
	Run: func(cmd *cobra.Command, args []string) {
		pidPath := getPIDFilePath()

		// Read PID file
		pidBytes, err := os.ReadFile(pidPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("⚠️  Daemon is not running (no PID file found)")
				return
			}
			log.Fatalf("Failed to read PID file: %v", err)
		}

		var pid int
		if _, err := fmt.Sscanf(string(pidBytes), "%d", &pid); err != nil {
			log.Fatalf("Invalid PID file: %v", err)
		}

		// Find process
		process, err := os.FindProcess(pid)
		if err != nil {
			log.Fatalf("Failed to find process: %v", err)
		}

		// Send SIGTERM for graceful shutdown
		fmt.Printf("🛑 Stopping daemon (PID: %d)...\n", pid)
		if err := process.Signal(syscall.SIGTERM); err != nil {
			if err == os.ErrProcessDone {
				fmt.Println("⚠️  Daemon was already stopped")
				os.Remove(pidPath)
				return
			}
			log.Fatalf("Failed to stop daemon: %v", err)
		}

		// Wait a moment and verify it stopped
		fmt.Println("   Waiting for graceful shutdown...")

		// Check if process is still running
		if err := process.Signal(os.Signal(nil)); err != nil {
			// Process is gone
			fmt.Println("✅ Daemon stopped successfully")
			os.Remove(pidPath)
			return
		}

		// If still running after a few seconds, could force kill
		fmt.Println("⚠️  Daemon may still be shutting down")
		fmt.Println("   Run 'greycli status' to verify")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
