package cmd

import (
	"fmt"
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/spf13/cobra"
)

// Flags for generate command
var (
	genModel  string
	genPrompt string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate text using an LLM",
	Long:  "Generates text completion from a prompt using Ollama LLM models.",
	Run: func(cmd *cobra.Command, args []string) {
		if genModel == "" || genPrompt == "" {
			log.Fatalf("Missing required flags: --model and --prompt are required.")
		}

		resp, err := api.Generate(genModel, genPrompt)
		if err != nil {
			log.Fatalf("Generation failed: %v", err)
		}

		fmt.Printf("✅ Text generated successfully!\n")
		fmt.Printf("Model: %s\n", genModel)
		fmt.Printf("Tokens: %d\n", resp.TokensGenerated)
		fmt.Printf("Latency: %dms\n", resp.LatencyMs)
		fmt.Printf("\n%s\n", resp.Response)
	},
}

func init() {
	generateCmd.Flags().StringVar(&genModel, "model", "llama3.2:3b", "LLM model to use for generation")
	generateCmd.Flags().StringVar(&genPrompt, "prompt", "", "Text prompt for generation (required)")

	rootCmd.AddCommand(generateCmd)
}
