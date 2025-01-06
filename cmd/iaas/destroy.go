package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy all resources from the cloud",
	Run: func(_ *cobra.Command, _ []string) {
		current, err := state.GetState()
		if err != nil {
			fmt.Printf("failed to get state: %v\n", err)
			return
		}

		for k, v := range *current {
			if err := resourceManager.DeleteResource(v.Region, k); err != nil {
				fmt.Printf("failed to delete resource: %v\n", err)
				return
			}

			delete(*current, k)
		}

		if err := state.WriteState(*current); err != nil {
			fmt.Printf("failed to write state: %v\n", err)
			return
		}

		fmt.Println("All Resources destroyed successfully")
	},
}
