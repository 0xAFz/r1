package iaas

import (
	"fmt"
	"sync"

	"github.com/0xAFz/r1/internal/api"
	"github.com/0xAFz/r1/internal/state"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy R1-managed infrastructure.",
	Run: func(_ *cobra.Command, _ []string) {
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(current) == 0 {
			fmt.Println("No objects need to be destroyed.")
			return
		}

		var wg sync.WaitGroup
		var mu sync.Mutex
		var nextState []api.ArvanInstanceResource

		for _, v := range current {
			resource := v
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := provider.DeleteInstance(resource.Region, resource.Data.ID); err != nil {
					fmt.Printf("arvancloud_compute_instance.%s: %v\n", resource.Data.Name, err)
					mu.Lock()
					nextState = append(nextState, resource)
					mu.Unlock()
					return
				}
				fmt.Printf("arvancloud_compute_instance.%s: Destruction complete\n", resource.Data.Name)

			}()
		}

		wg.Wait()

		if err := state.WriteCurrentState(nextState); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}
