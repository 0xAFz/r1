package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var (
	name          string
	count         int
	enableBackups bool
	flavorID      string
	sshKey        bool
	keyName       string
	diskSize      int
	enableIPv6    bool
	imageID       string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a vm on the cloud",
	Run: func(_ *cobra.Command, _ []string) {
		current, err := state.GetState()
		if err != nil {
			fmt.Printf("failed to get state: %v\n", err)
			return
		}

		vmRequest := api.CreateVMRequest{
			Name:         name,
			BackupName:   nil,
			Count:        count,
			EnableBackup: enableBackups,
			NetworkIDs:   []string{"b81f15e7-38ba-483e-8aa3-39ace15094d7"},
			FlavorID:     flavorID,
			SecurityGroupIDs: []map[string]any{
				{"name": "010c38b1-5259-44dd-a82c-161aa492809b"},
			},
			SSHKey:        sshKey,
			KeyName:       keyName,
			DiskSize:      diskSize,
			InitScript:    "",
			HAEnabled:     true,
			ServerVolumes: []map[string]any{},
			EnableIPv6:    enableBackups,
			ImageID:       imageID,
		}

		r, err := resourceManager.CreateResource(vmRequest)
		if err != nil {
			fmt.Println("failed to create resource:", err)
			return
		}

		(*current)[r.Data.ID] = struct {
			Status string   `json:"status"`
			IP     []string `json:"ip"`
		}{
			Status: "",
			IP:     []string{},
		}

		if err := state.WriteState(*current); err != nil {
			fmt.Println("ID: ", r.Data.ID)
			fmt.Println("failed to update state:", err)
			return
		}

		fmt.Println("Resource created successfully")
	},
}

func init() {
	createCmd.Flags().StringVar(&name, "name", "", "Name of the resource (required)")
	createCmd.Flags().IntVar(&count, "count", 1, "Number of resources to create (default: 1)")
	createCmd.Flags().BoolVar(&enableBackups, "enable-backups", false, "Enable backups (default: false)")
	createCmd.Flags().StringVar(&flavorID, "flavor-id", "", "ID of the flavor (required)")
	createCmd.Flags().BoolVar(&sshKey, "ssh-key", false, "Enable SSH key (default: false)")
	createCmd.Flags().StringVar(&keyName, "key-name", "kumo", "Name of the SSH key (default: kumo)")
	createCmd.Flags().IntVar(&diskSize, "disk-size", 50, "Size of the disk in GB (default: 50)")
	createCmd.Flags().BoolVar(&enableIPv6, "enable-ipv6", false, "Enable IPv6 (default: false)")
	createCmd.Flags().StringVar(&imageID, "image-id", "6ea4bc95-fd05-4695-9c0a-d5a566d6b9da", "ID of the image (default: debian 12)")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("flavor-id")
}
