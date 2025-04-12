package iaas

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/0xAFz/r1/internal/api"
	"github.com/0xAFz/r1/internal/state"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Creates infrastructure according to R1 configuration files in the current directory.",
	Run: func(_ *cobra.Command, _ []string) {
		desired, err := state.ReadDesiredState()
		if err != nil {
			log.Fatal(err)
			return
		}
		current, err := state.ReadCurrentState()
		if err != nil {
			log.Fatal(err)
			return
		}

		desiredMap := make(map[string]api.ArvanInstanceRequest)
		for _, req := range desired {
			desiredMap[req.Data.Name] = req
		}
		currentMap := make(map[string]api.ArvanInstanceResource)
		for _, vm := range current {
			currentMap[vm.Data.Name] = vm
		}

		var newState []api.ArvanInstanceResource

		var wg sync.WaitGroup

		for _, req := range desired {
			if existing, exists := currentMap[req.Data.Name]; !exists {
				wg.Add(1)
				// Create VM if it doesn’t exist or isn’t active
				go func() {
					defer wg.Done()
					fmt.Printf("arvancloud_compute_instance.%s: Creating...\n", req.Data.Name)
					resp, err := provider.CreateInstance(req)
					if err != nil {
						fmt.Printf("%s: %v\n", req.Data.Name, err)
						return
					}

					newResource := api.ArvanInstanceResource{
						Region:        req.Region,
						ArvanInstance: *resp,
					}

					start := time.Now()
					waitCount := 10
					for {
						time.Sleep(time.Second * 10)
						fmt.Printf("arvancloud_compute_instance.%s: Still creating... [%ds elapsed]\n", req.Data.Name, waitCount)
						waitCount += 10
						ins, err := provider.GetInstance(newResource.Region, newResource.Data.ID)
						if err != nil {
							fmt.Printf("arvancloud_compute_instance.%s: %v", req.Data.Name, err)
							continue
						}
						if ins.Data.Status != "ACTIVE" {
							continue
						}
						newResource.Data = ins.Data
						break
					}
					newState = append(newState, newResource)
					fmt.Printf("arvancloud_compute_instance.%s: Creation complete after %v\n", req.Data.Name, time.Since(start))
				}()
			} else {
				// Keep existing VM
				newState = append(newState, existing)
			}
		}

		// Process current state: destroy unwanted VMs
		for name, vm := range currentMap {
			if _, keep := desiredMap[name]; !keep {
				wg.Add(1)
				go func() {
					if err := provider.DeleteInstance(vm.Region, vm.Data.ID); err != nil {
						fmt.Printf("arvancloud_compute_instance.%s: %v", vm.Data.Name, err)
						return
					}
					fmt.Printf("arvancloud_compute_instance.%s: Destruction complete\n", vm.Data.Name)
				}()
			}
		}

		wg.Wait()

		if err := state.WriteCurrentState(newState); err != nil {
			log.Fatal(err)
			return
		}
	},
}
