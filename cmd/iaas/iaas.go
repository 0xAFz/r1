package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/config"
	"github.com/0xAFz/kumo/internal/vm"
	"github.com/spf13/cobra"
)

var resourceManager *vm.ResourceManager

var IaaSCmd = &cobra.Command{
	Use:   "iaas",
	Short: "Manage iaas actions",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiClient := api.NewAPIClient(vm.BaseURL, config.AppConfig.APIKey)
		resourceManager = vm.NewResourceManager(apiClient)
	},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Usage: kumo iaas <action> (apply|destroy|status)")
	},
}

func init() {
	config.LoadConfig()

	IaaSCmd.AddCommand(applyCmd)
	IaaSCmd.AddCommand(destroyCmd)
	IaaSCmd.AddCommand(statusCmd)
}
