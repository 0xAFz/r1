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

func (r *ResourceManager) CreateResource(req api.IaasCreateRequest) (*api.IaasResponse, error) {
	endpoint := fmt.Sprintf("/%s/servers", req.Region)

	resp, err := r.client.Post(endpoint, req.Data)
	if err != nil {
		return nil, fmt.Errorf("create resource: %w", err)
	}

	var createResp api.IaasResponse
	if err := json.Unmarshal(resp, &createResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &createResp, nil
}

func (r *ResourceManager) GetResource(region, id string) (*api.IaasResponse, error) {
	endpoint := fmt.Sprintf("/%s/servers/%s", region, id)

	resp, err := r.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get resource: %w", err)
	}

	var iaasResp api.IaasResponse
	if err := json.Unmarshal(resp, &iaasResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &iaasResp, nil
}

func (r *ResourceManager) DeleteResource(region, id string) error {
	endpoint := fmt.Sprintf("/%s/servers/%s?forceDelete=true", region, id)

	_, err := r.client.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("delete resource: %w", err)
	}

	return nil
}
