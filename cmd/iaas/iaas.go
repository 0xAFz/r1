package iaas

import (
	"fmt"

	"github.com/0xAFz/r1/internal/api"
	"github.com/0xAFz/r1/internal/config"
	"github.com/0xAFz/r1/internal/vm"
	"github.com/spf13/cobra"
)

var provider *vm.Provider

var IaaSCmd = &cobra.Command{
	Use:   "iaas",
	Short: "Manage iaas actions",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiClient := api.NewAPIClient(vm.BaseURL, config.AppConfig.ApiKey)
		provider = vm.NewProvider(apiClient)
	},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("action required")
	},
}

func init() {
	config.LoadConfig()

	IaaSCmd.AddCommand(applyCmd)
	IaaSCmd.AddCommand(destroyCmd)
}
