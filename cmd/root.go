package cmd

import (
	"os"

	"github.com/0xAFz/r1/cmd/iaas"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "r1",
	Short: "ArvanCloud IaaS CLI Tool",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(iaas.IaaSCmd)
	rootCmd.AddCommand(StateCmd)
}
