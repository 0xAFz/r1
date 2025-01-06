package iaas

import (
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var (
	region          string
	name            string
	count           int
	enableBackups   bool
	flavorID        string
	sshKey          bool
	keyName         string
	diskSize        int
	enableIPv6      bool
	imageID         string
	securityGroupID string
	networkIDs      []string
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
			NetworkIDs:   networkIDs,
			FlavorID:     flavorID,
			SecurityGroupIDs: []map[string]any{
				{"name": securityGroupID},
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

		r, err := resourceManager.CreateResource(region, vmRequest)
		if err != nil {
			fmt.Println("failed to create resource:", err)
			return
		}

		(*current)[r.Data.ID] = struct {
			Status string   `json:"status"`
			IP     []string `json:"ip"`
			Region string   `json:"region"`
		}{
			Status: "BUILD",
			IP:     []string{},
			Region: region,
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
	createCmd.Flags().IntVar(&count, "count", 1, "Number of resources to create")
	createCmd.Flags().BoolVar(&enableBackups, "enable-backups", false, "Enable backups")
	createCmd.Flags().StringVar(&flavorID, "flavor-id", "", "ID of the flavor (required)")
	createCmd.Flags().BoolVar(&sshKey, "ssh-key", false, "Enable SSH key")
	createCmd.Flags().StringVar(&keyName, "key-name", "kumo", "Name of the SSH key")
	createCmd.Flags().IntVar(&diskSize, "disk-size", 25, "Size of the disk in GB")
	createCmd.Flags().BoolVar(&enableIPv6, "enable-ipv6", false, "Enable IPv6")
	createCmd.Flags().StringVar(&imageID, "image-id", "", "ID of the image")
	createCmd.Flags().StringSliceVar(&networkIDs, "network-ids", nil, "ID of the network")
	createCmd.Flags().StringVar(&securityGroupID, "security-group-id", "", "ID of the security group")
	createCmd.Flags().StringVar(&region, "region", "", "Region of the resource (required)")

	createCmd.MarkFlagRequired("region")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("flavor-id")
	createCmd.MarkFlagRequired("image-id")
	createCmd.MarkFlagRequired("network-ids")
	createCmd.MarkFlagRequired("security-group-id")
}
