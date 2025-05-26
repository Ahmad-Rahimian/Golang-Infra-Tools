package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secret-decoder",
	Short: "Decode secrets from your Kubernetes cluster",
	Long:  `A kubectl plugin to decode Kubernetes Secrets in human-readable format from any namespace.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}