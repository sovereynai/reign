package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/spf13/cobra"
)

var creditsJSON bool

var creditsCmd = &cobra.Command{
	Use:   "credits",
	Short: "Show your current credit balance",
	Long:  "Fetches the current credit balance associated with your node or wallet from the Fist Network.",
	Run: func(cmd *cobra.Command, args []string) {
		credits, err := api.FetchCredits()
		if err != nil {
			log.Fatalf("Failed to fetch credits: %v", err)
		}

		if creditsJSON {
			data, _ := json.MarshalIndent(map[string]int{"credits": credits}, "", "  ")
			fmt.Println(string(data))
		} else {
			fmt.Printf("💰 Current Credits: %d\n", credits)
		}
	},
}

func init() {
	creditsCmd.Flags().BoolVar(&creditsJSON, "json", false, "Output credits in JSON format")
	rootCmd.AddCommand(creditsCmd)
}
