package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	daemonPath string
	background bool
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Knuckle daemon",
	Long:  "Starts the Knuckle daemon (greycore) in the background.",
	Run: func(cmd *cobra.Command, args []string) {
		// Determine daemon binary path
		if daemonPath == "" {
			// Try to find in PATH or standard locations
			daemonPath = findDaemonBinary()
			if daemonPath == "" {
				log.Fatalf("Could not find knuckle daemon. Use --daemon-path to specify location.")
			}
		}

		// Check if already running
		if isDaemonRunning() {
			fmt.Println("⚠️  Daemon is already running")
			fmt.Println("   Run 'greycli status' to check status")
			return
		}

		// Start the daemon
		fmt.Printf("🚀 Starting Knuckle daemon...\n")
		fmt.Printf("   Binary: %s\n", daemonPath)

		var startCmd *exec.Cmd
		if background {
			// Start in background
			startCmd = exec.Command(daemonPath, "serve")
			startCmd.Stdout = nil
			startCmd.Stderr = nil
			startCmd.Stdin = nil

			if err := startCmd.Start(); err != nil {
				log.Fatalf("Failed to start daemon: %v", err)
			}

			// Write PID file
			pidPath := getPIDFilePath()
			if err := os.WriteFile(pidPath, []byte(fmt.Sprintf("%d", startCmd.Process.Pid)), 0644); err != nil {
				log.Printf("Warning: Could not write PID file: %v", err)
			}

			fmt.Printf("✅ Daemon started successfully (PID: %d)\n", startCmd.Process.Pid)
			fmt.Printf("   PID file: %s\n", pidPath)
		} else {
			// Start in foreground
			startCmd = exec.Command(daemonPath, "serve")
			startCmd.Stdout = os.Stdout
			startCmd.Stderr = os.Stderr
			startCmd.Stdin = os.Stdin

			if err := startCmd.Run(); err != nil {
				log.Fatalf("Daemon exited with error: %v", err)
			}
		}

		fmt.Println("\n💡 Use 'greycli status' to check daemon status")
	},
}

func init() {
	startCmd.Flags().StringVar(&daemonPath, "daemon-path", "", "Path to knuckle daemon binary")
	startCmd.Flags().BoolVarP(&background, "background", "d", true, "Run daemon in background")

	rootCmd.AddCommand(startCmd)
}

// findDaemonBinary attempts to locate the knuckle daemon binary
func findDaemonBinary() string {
	// Check common locations
	candidates := []string{
		"knuckle",
		"greycore",
		"/usr/local/bin/greycore",
		"/usr/local/bin/knuckle",
		filepath.Join(os.Getenv("HOME"), ".local", "bin", "knuckle"),
	}

	// On Homebrew installations
	if runtime.GOOS == "darwin" {
		candidates = append(candidates,
			"/opt/homebrew/bin/greycore",
			"/usr/local/opt/greymattr/bin/greycore",
		)
	}

	for _, candidate := range candidates {
		if path, err := exec.LookPath(candidate); err == nil {
			return path
		}
	}

	return ""
}

// getPIDFilePath returns the path to the daemon PID file
func getPIDFilePath() string {
	if runtime.GOOS == "darwin" {
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "knuckle", "daemon.pid")
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "knuckle", "daemon.pid")
}

// isDaemonRunning checks if the daemon is currently running
func isDaemonRunning() bool {
	pidPath := getPIDFilePath()
	pidBytes, err := os.ReadFile(pidPath)
	if err != nil {
		return false
	}

	var pid int
	if _, err := fmt.Sscanf(string(pidBytes), "%d", &pid); err != nil {
		return false
	}

	// Check if process exists
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// On Unix, Signal 0 checks if process exists
	if err := process.Signal(os.Signal(nil)); err != nil {
		return false
	}

	return true
}
