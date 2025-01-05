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

		for k := range *current {
			r, err := resourceManager.GetResource(k)
			if err != nil {
				fmt.Printf("failed to get resource: %v\n", err)
				return
			}

			ips := make([]string, 0)

			for _, v := range r.Data.Addresses {
				for _, a := range v {
					ips = append(ips, a.Addr)
				}
			}

			(*current)[k] = struct {
				Status string   `json:"status"`
				IP     []string `json:"ip"`
			}{
				Status: r.Data.Status,
				IP:     ips,
			}
		}

		if err := state.WriteState(*current); err != nil {
			fmt.Printf("failed to write state: %v\n", err)
			return
		}
	},
}
