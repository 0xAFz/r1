package iaas

import (
	"fmt"

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
				r := api.IaasResource{
					Region:       req.Region,
					IaasResponse: *resp,
				}
				newState = append(newState, r)
				fmt.Printf("Created VM %s with ID %s\n", req.Data.Name, resp.Data.ID)
			} else {
				// Keep existing VM
				newState = append(newState, existing)
			}
		}

		// Process current state: delete unwanted VMs
		for name, vm := range currentMap {
			if _, keep := desiredMap[name]; !keep {
				err := resourceManager.DeleteResource(vm.Region, vm.Data.ID)
				if err != nil {
					fmt.Println(err)
					continue
				} else {
					fmt.Printf("Deleted Resource: %s\n", name)
				}
			}
		}

		if err := state.WriteCurrentState(newState); err != nil {
			fmt.Println("writing state:", err)
			return
		}
	},
}
