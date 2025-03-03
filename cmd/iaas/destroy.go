package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy Kumo-managed infrastructure.",
	Run: func(_ *cobra.Command, _ []string) {
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range current {
			if err := resourceManager.DeleteResource(v.Region, v.Data.ID); err != nil {
				fmt.Printf("%s: %v\n", v.Data.Name, err)
				return
			}
			fmt.Printf("Deleted Resource: %s\n", v.Data.Name)
		}

		if err := state.WriteCurrentState([]api.IaasResource{}); err != nil {
			fmt.Println("writing current state:", err)
			return
		}
	},
}
