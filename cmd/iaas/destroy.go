package iaas

import (
	"fmt"
	"sync"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy kumo-managed infrastructure.",
	Run: func(_ *cobra.Command, _ []string) {
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		var wg sync.WaitGroup

		for _, v := range current {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := provider.DeleteInstance(v.Region, v.Data.ID); err != nil {
					fmt.Printf("arvancloud_compute_instance.%s: %v\n", v.Data.Name, err)
					return
				}
				fmt.Printf("arvancloud_compute_instance.%s: Destruction complete\n", v.Data.Name)
			}()
		}

		wg.Wait()

		if err := state.WriteCurrentState([]api.ArvanInstanceResource{}); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}
