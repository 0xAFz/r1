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

		var wg sync.WaitGroup
		removeChan := make(chan int, len(current))

		for i := range current {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if err := provider.DeleteInstance(current[i].Region, current[i].Data.ID); err != nil {
					fmt.Printf("arvancloud_compute_instance.%s: %v\n", current[i].Data.Name, err)
					return
				}
				fmt.Printf("arvancloud_compute_instance.%s: Destruction complete\n", current[i].Data.Name)
				removeChan <- i
			}(i)
		}

		var removeIndexes []int
		go func() {
			wg.Wait()
			close(removeChan)
		}()

		for i := range removeChan {
			removeIndexes = append(removeIndexes, i)
		}

		for i := len(removeIndexes) - 1; i >= 0; i-- {
			current = removeResource(current, removeIndexes[i])
		}

		if err := state.WriteCurrentState(current); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}

func removeResource(s []api.ArvanInstanceResource, i int) []api.ArvanInstanceResource {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
