package vm

import (
	"encoding/json"
	"fmt"

	"github.com/0xAFz/kumo/internal/api"
)

const (
	BaseURL = "https://napi.arvancloud.ir/ecc/v1/regions"
)

type Provider struct {
	client *api.APIClient
}

func NewProvider(client *api.APIClient) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) CreateInstance(req api.ArvanInstanceRequest) (*api.ArvanInstance, error) {
	endpoint := fmt.Sprintf("/%s/servers", req.Region)
	resp, err := p.client.Post(endpoint, req.Data)
	if err != nil {
		return nil, fmt.Errorf("create instance: %w", err)
	}
	var instance api.ArvanInstance
	if err := json.Unmarshal(resp, &instance); err != nil {
		return nil, fmt.Errorf("unmarshal instance: %w", err)
	}
	return &instance, nil
}

func (p *Provider) GetInstance(region, id string) (*api.ArvanInstance, error) {
	endpoint := fmt.Sprintf("/%s/servers/%s", region, id)
	resp, err := p.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get instance: %w", err)
	}
	var instance api.ArvanInstance
	if err := json.Unmarshal(resp, &instance); err != nil {
		return nil, fmt.Errorf("unmarshal instance: %w", err)
	}
	return &instance, nil
}

func (p *Provider) DeleteInstance(region, id string) error {
	endpoint := fmt.Sprintf("/%s/servers/%s?forceDelete=true", region, id)
	_, err := p.client.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("delete instance: %w", err)
	}
	return nil
}
