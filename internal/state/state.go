package state

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	filename = "state.json"
)

type State map[string]struct {
	Status string   `json:"status"`
	IP     []string `json:"ip"`
}

func WriteState(state State) error {
	file, err := json.MarshalIndent(state, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0o644)
}

func GetState() (*State, error) {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			if err := WriteState(State{}); err != nil {
				return nil, fmt.Errorf("failed to write initial state: %v", err)
			}
			return &State{}, nil
		}
		return nil, fmt.Errorf("error checking state file status: %v", err)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var state State
	if err := json.Unmarshal(file, &state); err != nil {
		return nil, err
	}

	return &state, nil
}
