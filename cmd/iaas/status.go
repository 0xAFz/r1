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

		for k, v := range *current {
			r, err := resourceManager.GetResource(v.Region, k)
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

			v.IP = ips
			v.Status = r.Data.Status

			(*current)[k] = v
		}

		if err := state.WriteState(*current); err != nil {
			fmt.Printf("failed to write state: %v\n", err)
			return
		}
	},
}
