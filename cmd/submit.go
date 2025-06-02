package cmd

import (
	"fmt"
	"log"

	"github.com/Leathal1/greycli/internal/api"
	"github.com/spf13/cobra"
)

// Flags
var (
	model      string
	imagePath  string
	redundancy int
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit an image for inference",
	Long:  "Submits an image to the Fist Network inference engine with a selected model and optional redundancy.",
	Run: func(cmd *cobra.Command, args []string) {
		if model == "" || imagePath == "" {
			log.Fatalf("Missing required flags: --model and --image are required.")
		}

       // note: SubmitJob takes (imagePath, model, redundancy)
		jobResp, err := api.SubmitJob(imagePath, model, redundancy)
		if err != nil {
			log.Fatalf("Submission failed: %v", err)
		}

		fmt.Printf("✅ Job submitted successfully!\nJob ID: %s\nModel: %s\nImage: %s\nRedundancy: %d\n", jobResp.JobID, model, imagePath, redundancy)
	},
}

func init() {
	submitCmd.Flags().StringVar(&model, "model", "", "Model to use for inference (required)")
	submitCmd.Flags().StringVar(&imagePath, "image", "", "Path to the input image file (required)")
	submitCmd.Flags().IntVar(&redundancy, "redundancy", 1, "Number of redundant peers to run inference (optional)")

	rootCmd.AddCommand(submitCmd)
}
