package bootstrap

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Setup performs first-run initialization
func Setup() error {
	if !isFirstRun() {
		return nil // Already set up
	}

	fmt.Println()
	fmt.Println("🚀 Welcome to Sovereyn!")
	fmt.Println()
	fmt.Println("⏳ First-time setup (this will only happen once)...")
	fmt.Println()

	// Check system requirements
	if err := checkSystemRequirements(); err != nil {
		return fmt.Errorf("system requirements not met: %w", err)
	}
	fmt.Println("   ✅ System requirements met")

	// Ensure Ollama
	if err := ensureOllama(); err != nil {
		return fmt.Errorf("ollama setup failed: %w", err)
	}

	// Check if throne is installed (optional - not required if user points to remote daemon)
	if isThroneInstalled() {
		fmt.Println("   ✅ Throne daemon installed")
	} else {
		fmt.Println("   ℹ️  Throne binary not in PATH (will use THRONE_URL if set)")
	}

	// Pull default model
	if err := ensureDefaultModel("llama3.2:3b"); err != nil {
		// Not fatal - user can pull models later
		fmt.Printf("   ⚠️  Could not pull default model: %v\n", err)
		fmt.Println("      You can pull models later with: ollama pull llama3.2:3b")
	}

	fmt.Println()
	fmt.Println("✨ Setup complete!")
	fmt.Println()

	// Mark as complete
	markSetupComplete()

	// Offer to start throne
	fmt.Println("📋 Next steps:")
	fmt.Println("   1. Start throne daemon: throne serve &")
	fmt.Println("   2. Run your first inference: reign chat \"Hello world\"")
	fmt.Println()

	return nil
}

// EnsureThroneRunning checks if throne is running and offers to start it
func EnsureThroneRunning() error {
	// Check if throne is responding
	if isThroneResponding() {
		return nil
	}

	fmt.Println()
	fmt.Println("⚠️  Throne daemon not running!")
	fmt.Println()
	fmt.Println("Start it with:")
	fmt.Println("   throne serve &")
	fmt.Println()
	fmt.Println("Or run in foreground:")
	fmt.Println("   throne serve")
	fmt.Println()

	return fmt.Errorf("throne daemon not running")
}

func isFirstRun() bool {
	sovereignHome := getSovereignHome()
	setupMarker := filepath.Join(sovereignHome, ".setup_complete")
	_, err := os.Stat(setupMarker)
	return os.IsNotExist(err)
}

func markSetupComplete() {
	sovereignHome := getSovereignHome()
	os.MkdirAll(sovereignHome, 0755)
	setupMarker := filepath.Join(sovereignHome, ".setup_complete")
	os.WriteFile(setupMarker, []byte(time.Now().Format(time.RFC3339)), 0644)
}

func getSovereignHome() string {
	if home := os.Getenv("SOVEREYN_HOME"); home != "" {
		return home
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".sovereyn")
}

func checkSystemRequirements() error {
	// Check architecture
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		return fmt.Errorf("unsupported architecture: %s (need amd64 or arm64)", runtime.GOARCH)
	}

	// Check OS
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" && runtime.GOOS != "windows" {
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// TODO: Check RAM (4GB minimum)
	// TODO: Check disk space (10GB minimum for models)

	return nil
}

func ensureOllama() error {
	// Check if already installed
	if isOllamaInstalled() {
		fmt.Println("   ✅ Ollama already installed")

		// Make sure it's running
		if !isOllamaRunning() {
			fmt.Println("   ⏳ Starting Ollama service...")
			if err := startOllama(); err != nil {
				fmt.Printf("   ⚠️  Could not start Ollama automatically: %v\n", err)
				fmt.Println("      Please run: ollama serve &")
			} else {
				time.Sleep(2 * time.Second) // Wait for startup
				fmt.Println("   ✅ Ollama service started")
			}
		} else {
			fmt.Println("   ✅ Ollama service running")
		}
		return nil
	}

	// Not installed - install it
	fmt.Println("   ⏳ Installing Ollama...")
	return installOllama()
}

func isOllamaInstalled() bool {
	_, err := exec.LookPath("ollama")
	return err == nil
}

func isOllamaRunning() bool {
	cmd := exec.Command("ollama", "list")
	err := cmd.Run()
	return err == nil
}

func startOllama() error {
	cmd := exec.Command("ollama", "serve")
	if err := cmd.Start(); err != nil {
		return err
	}
	// Detach from process
	return nil
}

func installOllama() error {
	switch runtime.GOOS {
	case "darwin":
		return installOllamaMacOS()
	case "linux":
		return installOllamaLinux()
	case "windows":
		return installOllamaWindows()
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

func installOllamaMacOS() error {
	// Check for Homebrew
	if _, err := exec.LookPath("brew"); err == nil {
		fmt.Println("      • Using Homebrew...")
		cmd := exec.Command("brew", "install", "ollama")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("brew install failed: %w", err)
		}
		fmt.Println("   ✅ Ollama installed via Homebrew")
		return startOllama()
	}

	// No homebrew - give manual instructions
	fmt.Println()
	fmt.Println("   ⚠️  Homebrew not found")
	fmt.Println()
	fmt.Println("   Install Ollama manually:")
	fmt.Println("      1. Visit: https://ollama.ai/download")
	fmt.Println("      2. Download and run the installer")
	fmt.Println("      3. Run: ollama serve &")
	fmt.Println()
	return fmt.Errorf("manual installation required")
}

func installOllamaLinux() error {
	fmt.Println("      • Downloading installer...")

	// Use the official install script
	cmd := exec.Command("curl", "-fsSL", "https://ollama.ai/install.sh")
	script, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to download installer: %w", err)
	}

	// Run the install script
	cmd = exec.Command("sh", "-c", string(script))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}

	fmt.Println("   ✅ Ollama installed")
	return startOllama()
}

func installOllamaWindows() error {
	fmt.Println()
	fmt.Println("   ⚠️  Automatic installation not available on Windows")
	fmt.Println()
	fmt.Println("   Install Ollama manually:")
	fmt.Println("      1. Visit: https://ollama.ai/download")
	fmt.Println("      2. Download OllamaSetup.exe")
	fmt.Println("      3. Run the installer")
	fmt.Println("      4. Ollama will start automatically")
	fmt.Println()
	return fmt.Errorf("manual installation required")
}

func ensureDefaultModel(model string) error {
	// Check if model exists
	cmd := exec.Command("ollama", "list")
	output, err := cmd.Output()
	if err == nil && strings.Contains(string(output), model) {
		fmt.Printf("   ✅ Model '%s' already available\n", model)
		return nil
	}

	// Pull the model
	fmt.Printf("   ⏳ Pulling default model (%s)...\n", model)
	fmt.Println("      (This may take 2-5 minutes for first download)")
	fmt.Println()

	cmd = exec.Command("ollama", "pull", model)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to pull model: %w", err)
	}

	fmt.Println()
	fmt.Printf("   ✅ Model '%s' ready!\n", model)
	return nil
}

func isThroneInstalled() bool {
	_, err := exec.LookPath("throne")
	return err == nil
}

func isThroneResponding() bool {
	// Check if THRONE_URL is explicitly set
	if throneURL := os.Getenv("THRONE_URL"); throneURL != "" {
		cmd := exec.Command("curl", "-s", "-o", "/dev/null", "-w", "%{http_code}",
			throneURL+"/healthz")
		output, err := cmd.Output()
		if err == nil && string(output) == "200" {
			return true
		}
		return false
	}

	// Try common ports
	ports := []string{"8080", "8081", "8082", "8083", "8090", "8091"}
	for _, port := range ports {
		cmd := exec.Command("curl", "-s", "-o", "/dev/null", "-w", "%{http_code}",
			fmt.Sprintf("http://localhost:%s/healthz", port))
		output, err := cmd.Output()
		if err == nil && string(output) == "200" {
			return true
		}
	}
	return false
}

// GetModelSize returns estimated model size for display
func GetModelSize(model string) string {
	sizes := map[string]string{
		"llama3.2:1b":     "1.3 GB",
		"llama3.2:3b":     "2.0 GB",
		"llama3.2:latest": "2.0 GB",
		"llama3.1:8b":     "4.7 GB",
		"llama3.1:70b":    "40 GB",
		"qwen2.5:3b":      "2.3 GB",
		"qwen2.5:7b":      "4.7 GB",
		"deepseek-r1:7b":  "4.7 GB",
	}
	if size, ok := sizes[model]; ok {
		return size
	}
	return "Unknown size"
}
