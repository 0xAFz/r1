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
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Usage: kumo iaas <action> (create|remove|status)")
	},
}

func init() {
	config.LoadConfig()

	apiClient := api.NewAPIClient(config.AppConfig.BaseURL, config.AppConfig.APIKey)

	resourceManager = vm.NewResourceManager(apiClient)

	IaaSCmd.AddCommand(createCmd)
	IaaSCmd.AddCommand(destroyCmd)
	IaaSCmd.AddCommand(statusCmd)
}
