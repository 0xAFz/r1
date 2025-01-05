package iaas

import (
	"fmt"
	"log"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/config"
	"github.com/0xAFz/kumo/internal/vm"
	"github.com/spf13/cobra"
)

var (
	region          string
	resourceManager *vm.ResourceManager
)

var IaaSCmd = &cobra.Command{
	Use:   "iaas",
	Short: "Manage iaas actions",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		baseURL := vm.BaseURL + "/" + region
		log.Println("baseURL: ", baseURL)

		apiClient := api.NewAPIClient(baseURL, config.AppConfig.APIKey)
		resourceManager = vm.NewResourceManager(apiClient)
	},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Usage: kumo iaas <action> (create|remove|status)")
	},
}

func init() {
	IaaSCmd.PersistentFlags().StringVar(&region, "region", "", "Region of the resource (required)")
	IaaSCmd.MarkPersistentFlagRequired("region")

	config.LoadConfig()

	IaaSCmd.AddCommand(createCmd)
	IaaSCmd.AddCommand(destroyCmd)
	IaaSCmd.AddCommand(statusCmd)
}
