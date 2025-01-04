package cmd

import (
	"os"

	"github.com/0xAFz/kumo/cmd/iaas"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Kumo",
	Short: "Kumo is a tool for working with Cloud API's",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(iaas.IaaSCmd)
}
