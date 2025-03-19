package iaas

import (
	"fmt"
	"time"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Creates infrastructure according to Kumo configuration files in the current directory.",
	Run: func(_ *cobra.Command, _ []string) {
		desired, err := state.ReadDesiredState()
		if err != nil {
			fmt.Println(err)
			return
		}
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		desiredMap := make(map[string]api.IaasCreateRequest)
		for _, req := range desired {
			desiredMap[req.Data.Name] = req
		}
		currentMap := make(map[string]api.IaasResource)
		for _, vm := range current {
			currentMap[vm.Data.Name] = vm
		}

		var newState []api.IaasResource

		for _, req := range desired {
			if existing, exists := currentMap[req.Data.Name]; !exists || existing.Data.Status != "ACTIVE" {
				// Create VM if it doesn’t exist or isn’t active
				resp, err := resourceManager.CreateResource(req)
				if err != nil {
					fmt.Printf("%s: %v\n", req.Data.Name, err)
					continue
				}
				newResource := api.IaasResource{
					Region:       req.Region,
					IaasResponse: *resp,
				}
				start := time.Now()
				waitCount := 1
				for {
					fmt.Printf("Waiting to create [%s] resource %d\n", req.Data.Name, waitCount)
					time.Sleep(time.Second * 1)
					waitCount++
					r, err := resourceManager.GetResource(req.Region, resp.Data.ID)
					if err != nil {
						fmt.Printf("failed to get resource: %v\n", err)
						continue
					}
					if r.Data.Status != "ACTIVE" {
						continue
					}
					newResource.Data = r.Data
					break
				}
				fmt.Printf("Created: %s - %v\n", req.Data.Name, time.Since(start))
				newState = append(newState, newResource)
			} else {
				// Keep existing VM
				newState = append(newState, existing)
			}
		}

		// Process current state: delete unwanted VMs
		for name, vm := range currentMap {
			if _, keep := desiredMap[name]; !keep {
				if err := resourceManager.DeleteResource(vm.Region, vm.Data.ID); err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("Destroyed: %s\n", name)
			}
		}

		if err := state.WriteCurrentState(newState); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}
