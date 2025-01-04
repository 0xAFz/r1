package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a vm on the cloud",
	Run: func(_ *cobra.Command, _ []string) {
		vmRequest := api.CreateVMRequest{
			Name:         "kumo",
			BackupName:   nil,
			Count:        1,
			EnableBackup: false,
			NetworkIDs:   []string{"b81f15e7-38ba-483e-8aa3-39ace15094d7"},
			FlaverID:     "sb1-1-1-0",
			SecurityGroupIDs: []map[string]any{
				{"name": "010c38b1-5259-44dd-a82c-161aa492809b"},
			},
			SSHKey:        true,
			KeyName:       "kumo",
			DiskSize:      15,
			InitScript:    "",
			HAEnabled:     true,
			ServerVolumes: []map[string]any{},
			EnableIPv6:    false,
			ImageID:       "6ea4bc95-fd05-4695-9c0a-d5a566d6b9da",
		}

		r, err := resourceManager.CreateResource(vmRequest)
		if err != nil {
			fmt.Println("failed to create resource:", err)
			return
		}

		s := state.State{
			ID: r.Data.ID,
		}

		if err := state.WriteState(s); err != nil {
			fmt.Println("ID: ", r.Data.ID)
			fmt.Println("failed to update state:", err)
			return
		}

		fmt.Println("Resource created successfully")
	},
}
