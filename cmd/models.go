package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/spf13/cobra"
)

var jsonOutput bool

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "List available models",
	Long:  "Fetches and displays the list of available models from the Fist Network tracker or local node.",
	Run: func(cmd *cobra.Command, args []string) {
		models, err := api.FetchModels()
		if err != nil {
			log.Fatalf("Failed to fetch models: %v", err)
		}

		if jsonOutput {
			data, _ := json.MarshalIndent(models, "", "  ")
			fmt.Println(string(data))
			return
		}

		fmt.Println("🧠 Available models:")
		for _, m := range models {
			fmt.Printf("- %s\n", m)
		}
	},
}

func init() {
	modelsCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output models in JSON format")
	rootCmd.AddCommand(modelsCmd)
}
