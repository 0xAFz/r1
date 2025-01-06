package vm

import (
	"encoding/json"
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
)

const (
	BaseURL = "https://napi.arvancloud.ir/ecc/v1/regions"
)

type ResourceManager struct {
	client *api.APIClient
}

func NewResourceManager(client *api.APIClient) *ResourceManager {
	return &ResourceManager{
		client: client,
	}
}

func (r *ResourceManager) CreateResource(region string, vmRequest api.CreateVMRequest) (*api.CreateVMResponse, error) {
	endpoint := fmt.Sprintf("/%s/servers", region)

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

func (r *ResourceManager) GetResource(region, id string) (*api.ResourceResponse, error) {
	endpoint := fmt.Sprintf("/%s/servers/%s", region, id)

	resp, err := r.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get resource: %w", err)
	}

	var resourceResponse api.ResourceResponse
	if err := json.Unmarshal(resp, &resourceResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resourceResponse, nil
}

func (r *ResourceManager) DeleteResource(region, id string) error {
	endpoint := fmt.Sprintf("/%s/servers/%s?forceDelete=true", region, id)

	_, err := r.client.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	return nil
}
