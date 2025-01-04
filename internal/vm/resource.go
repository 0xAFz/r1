package vm

import (
	"encoding/json"
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
)

type ResourceManager struct {
	client *api.APIClient
}

func NewResourceManager(client *api.APIClient) *ResourceManager {
	return &ResourceManager{
		client: client,
	}
}

func (r *ResourceManager) CreateResource(vmRequest api.CreateVMRequest) (*api.CreateVMResponse, error) {
	endpoint := "/servers"

	resp, err := r.client.Post(endpoint, vmRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	var createVMResponse api.CreateVMResponse
	if err := json.Unmarshal(resp, &createVMResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &createVMResponse, nil
}

func (r *ResourceManager) GetResource(id string) (*api.ResourceResponse, error) {
	endpoint := fmt.Sprintf("/servers/%s", id)

	resp, err := r.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	var resourceResponse api.ResourceResponse
	if err := json.Unmarshal(resp, &resourceResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resourceResponse, nil
}

func (r *ResourceManager) DeleteResource(id string) error {
	endpoint := fmt.Sprintf("/servers/%s?forceDelete=true", id)

	_, err := r.client.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	return nil
}
