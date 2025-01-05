package api

type CreateVMRequest struct {
	Name             string           `json:"name"`
	BackupName       *string          `json:"backup_name"`
	Count            int              `json:"count"`
	EnableBackup     bool             `json:"enable_backup"`
	NetworkIDs       []string         `json:"network_ids"`
	FlavorID         string           `json:"flavor_id"`
	SecurityGroupIDs []map[string]any `json:"security_groups"`
	SSHKey           bool             `json:"ssh_key"`
	KeyName          string           `json:"key_name"`
	DiskSize         int              `json:"disk_size"`
	InitScript       string           `json:"init_script"`
	HAEnabled        bool             `json:"ha_enabled"`
	ServerVolumes    []map[string]any `json:"server_volumes"`
	EnableIPv6       bool             `json:"enable_ipv6"`
	ImageID          string           `json:"image_id"`
}

// type Image struct {
// 	ID       string `json:"id"`
// 	Name     string `json:"name"`
// 	OS       string `json:"os"`
// 	OSVersion string `json:"os_version"`
// 	Status   string `json:"status"`
// 	Username string `json:"username"`
// 	Metadata struct {
// 		OSType  string `json:"os_type"`
// 		Username string `json:"username"`
// 	} `json:"metadata"`
// 	Documents []interface{} `json:"documents"`
// }

type CreateVMResponse struct {
	Data struct {
		ID string `json:"id"`
		// Name           string      `json:"name"`
		// Flavor         *string     `json:"flavor"`
		// Status         string      `json:"status"`
		// Image          Image       `json:"image"`
		// Created        string      `json:"created"`
		// Password       string      `json:"password"`
		// TaskState      *string     `json:"task_state"`
		// KeyName        string      `json:"key_name"`
		// ArNext         *string     `json:"ar_next"`
		// SecurityGroups interface{} `json:"security_groups"`
		// Addresses      interface{} `json:"addresses"`
		// Tags           []string    `json:"tags"`
		// HAEnabled      bool        `json:"ha_enabled"`
		// ClusterID      string      `json:"cluster_id"`
	} `json:"data"`
	// Message string `json:"message"`
}

type ResourceResponse struct {
	Data struct {
		ID string `json:"id"`
		// Name             string        `json:"name"`
		// Flavor           Flavor        `json:"flavor"`
		Status string `json:"status"`
		// Image            Image         `json:"image"`
		// Created          string        `json:"created"`
		// Password         string        `json:"password"`
		// TaskState        *string   `json:"task_state"`
		// KeyName          string        `json:"key_name"`
		// ARNext           string        `json:"ar_next"`
		// SecurityGroups   []SecurityGroup `json:"security_groups"`
		Addresses map[string][]Address `json:"addresses"`
		// Tags             []string      `json:"tags"`
		// HAEnabled        bool          `json:"ha_enabled"`
		// ClusterID        string        `json:"cluster_id"`
	} `json:"data"`
}

// type Flavor struct {
// 	ID     string `json:"id"`
// 	Name   string `json:"name"`
// 	RAM    int    `json:"ram"`
// 	Swap   string `json:"swap"`
// 	VCPUs  int    `json:"vcpus"`
// 	Disk   int    `json:"disk"`
// }

// type Image struct {
// 	ID       string            `json:"id"`
// 	Name     string            `json:"name"`
// 	OS       string            `json:"os"`
// 	OSVersion string           `json:"os_version"`
// 	Status   string            `json:"status"`
// 	Username string            `json:"username"`
// 	Metadata map[string]string `json:"metadata"`
// 	Documents []interface{}    `json:"documents"`
// }

// type SecurityGroup struct {
// 	ID          string   `json:"id"`
// 	Description string   `json:"description"`
// 	Name        string   `json:"name"`
// 	ReadOnly    bool     `json:"readonly"`
// 	Default     bool     `json:"default"`
// 	RealName    string   `json:"real_name"`
// 	Rules       interface{} `json:"rules"`
// 	IPAddresses []string `json:"ip_addresses"`
// }

type Address struct {
	// MacAddr  string `json:"mac_addr"`
	// Version  string `json:"version"`
	Addr string `json:"addr"`
	// Type     string `json:"type"`
	// IsPublic bool   `json:"is_public"`
}
