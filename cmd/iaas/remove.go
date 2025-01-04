package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove vm from the cloud",
	Run: func(_ *cobra.Command, _ []string) {
		s, err := state.GetState()
		if err != nil {
			fmt.Printf("failed to get state: %v\n", err)
			return
		}

		if err := resourceManager.DeleteResource(s.ID); err != nil {
			fmt.Printf("failed to delete resource: %v\n", err)
			return
		}

		if err := state.WriteState(state.State{}); err != nil {
			fmt.Printf("failed to update state: %v\n", err)
			return
		}

		fmt.Println("Resource removed successfully")
	},
}
