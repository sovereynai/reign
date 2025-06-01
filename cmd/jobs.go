package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/spf13/cobra"
)

var jobsJSON bool

var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Show recent inference jobs",
	Long:  "Retrieves a list of recently submitted inference jobs from the local node or network.",
	Run: func(cmd *cobra.Command, args []string) {
		jobs, err := api.FetchJobs()
		if err != nil {
			log.Fatalf("Failed to fetch jobs: %v", err)
		}

		if jobsJSON {
			data, _ := json.MarshalIndent(jobs, "", "  ")
			fmt.Println(string(data))
			return
		}

		fmt.Println("🧾 Recent Jobs:")
		for _, job := range jobs {
			fmt.Printf("- ID: %s | Model: %s | Status: %s\n", job.ID, job.Model, job.Status)
		}
	},
}

func init() {
	jobsCmd.Flags().BoolVar(&jobsJSON, "json", false, "Output jobs in JSON format")
	rootCmd.AddCommand(jobsCmd)
}
