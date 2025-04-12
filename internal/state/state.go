package state

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xAFz/kumo/internal/api"
)

const (
	current = ".state.json"
	desired = "kumo.json"
)

func ReadDesiredState() ([]api.ArvanInstanceRequest, error) {
	data, err := os.ReadFile(desired)
	if err != nil {
		return nil, fmt.Errorf("reading desired state: %v", err)
	}
	var reqs []api.ArvanInstanceRequest
	err = json.Unmarshal(data, &reqs)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling desired state: %v", err)
	}
	return reqs, nil
}

func ReadCurrentState() ([]api.ArvanInstanceResource, error) {
	data, err := os.ReadFile(current)
	if os.IsNotExist(err) {
		return []api.ArvanInstanceResource{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading current state: %v", err)
	}
	var state []api.ArvanInstanceResource
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling current state: %v", err)
	}
	return state, nil
}

func WriteCurrentState(state []api.ArvanInstanceResource) error {
	file, err := json.MarshalIndent(state, "", "    ")
	if err != nil {
		return fmt.Errorf("marshaling current state: %v", err)
	}
	return os.WriteFile(current, file, 0o644)
}
