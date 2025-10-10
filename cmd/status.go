package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check daemon status",
	Long:  "Checks if the Knuckle daemon is running and responds to health checks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Checking Knuckle daemon status...")
		fmt.Println()

		// Check PID file
		pidPath := getPIDFilePath()
		pidBytes, err := os.ReadFile(pidPath)

		var pid int
		var processRunning bool

		if err == nil {
			fmt.Sscanf(string(pidBytes), "%d", &pid)
			fmt.Printf("📄 PID File: %s\n", pidPath)
			fmt.Printf("   PID: %d\n", pid)

			// Check if process exists
			process, err := os.FindProcess(pid)
			if err == nil {
				if err := process.Signal(os.Signal(nil)); err == nil {
					processRunning = true
					fmt.Println("   Status: ✅ Process running")
				} else {
					fmt.Println("   Status: ❌ Process not found (stale PID file)")
				}
			} else {
				fmt.Println("   Status: ❌ Process not found")
			}
		} else {
			fmt.Printf("📄 PID File: Not found\n")
			fmt.Println("   Status: ⚠️  Daemon not started via greycli")
		}

		fmt.Println()

		// Check HTTP health endpoint
		fmt.Println("🌐 Health Check: http://localhost:8080/healthz")

		client := &http.Client{
			Timeout: 2 * time.Second,
		}

		resp, err := client.Get("http://localhost:8080/healthz")
		if err != nil {
			fmt.Printf("   Status: ❌ Not responding (%v)\n", err)
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("📊 OVERALL STATUS: ❌ OFFLINE")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			if processRunning {
				fmt.Println("\n⚠️  Process is running but not responding to HTTP")
				fmt.Println("   The daemon may be starting up or misconfigured")
			} else {
				fmt.Println("\n💡 Start the daemon with: greycli start")
			}
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("   Status: ✅ Healthy")

			// Get version header if available
			if version := resp.Header.Get("X-Version"); version != "" {
				fmt.Printf("   Version: %s\n", version)
			}
		} else {
			fmt.Printf("   Status: ⚠️  Unexpected status code: %d\n", resp.StatusCode)
		}

		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("📊 OVERALL STATUS: ✅ ONLINE")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		fmt.Println("💡 Available endpoints:")
		fmt.Println("   • Health:    http://localhost:8080/healthz")
		fmt.Println("   • Predict:   http://localhost:8080/predict")
		fmt.Println("   • Generate:  http://localhost:8080/generate")
		fmt.Println("   • Chat:      http://localhost:8080/chat")
		fmt.Println("   • Models:    http://localhost:8080/models")
		fmt.Println("   • Jobs:      http://localhost:8080/jobs")
		fmt.Println("   • Credits:   http://localhost:8080/credits")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
