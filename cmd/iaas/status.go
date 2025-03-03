package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get Kumo-managed infrastructure status.",
	Run: func(_ *cobra.Command, args []string) {
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, v := range current {
			r, err := resourceManager.GetResource(v.Region, v.Data.ID)
			if err != nil {
				fmt.Printf("failed to get resource: %v\n", err)
				return
			}
			current[i].Data = r.Data
		}

		if err := state.WriteCurrentState(current); err != nil {
			fmt.Printf("error writing state: %v\n", err)
			return
		}
	},
}
