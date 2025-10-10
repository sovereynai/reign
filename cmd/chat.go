package cmd

import (
	"fmt"
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/spf13/cobra"
)

// Flags for chat command
var (
	chatModel   string
	chatMessage string
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with an LLM",
	Long:  "Send a message to an LLM and receive a conversational response using Ollama models.",
	Run: func(cmd *cobra.Command, args []string) {
		if chatModel == "" || chatMessage == "" {
			log.Fatalf("Missing required flags: --model and --message are required.")
		}

		messages := []api.ChatMessage{
			{
				Role:    "user",
				Content: chatMessage,
			},
		}

		resp, err := api.Chat(chatModel, messages)
		if err != nil {
			log.Fatalf("Chat failed: %v", err)
		}

		fmt.Printf("✅ Chat response received!\n")
		fmt.Printf("Model: %s\n", resp.Model)
		fmt.Printf("Tokens: %d\n", resp.TokensGenerated)
		fmt.Printf("Latency: %dms\n", resp.LatencyMs)
		fmt.Printf("\n🤖 Assistant: %s\n", resp.Message.Content)
	},
}

func init() {
	chatCmd.Flags().StringVar(&chatModel, "model", "llama3.2:3b", "LLM model to use for chat")
	chatCmd.Flags().StringVar(&chatMessage, "message", "", "User message to send (required)")

	rootCmd.AddCommand(chatCmd)
}
