package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of a vm on the cloud",
	Run: func(_ *cobra.Command, args []string) {
		current, err := state.GetState()
		if err != nil {
			fmt.Printf("failed to get state: %v\n", err)
			return
		}

		r, err := resourceManager.GetResource(current.ID)
		if err != nil {
			fmt.Printf("failed to get resource: %v\n", err)
			return
		}

		ips := make([]string, 1)

		for _, v := range r.Data.Addresses {
			for _, a := range v {
				ips = append(ips, a.Addr)
			}
		}

		s := state.State{
			ID:     r.Data.ID,
			IP:     ips[0],
			Status: r.Data.Status,
		}

		fmt.Println(s)

		if err := state.WriteState(s); err != nil {
			fmt.Printf("failed to write state: %v\n", err)
			return
		}
	},
}
