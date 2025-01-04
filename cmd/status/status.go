package status

import (
	"encoding/json"
	"fmt"

	"github.com/0xAFz/kumo/internal/state"
	"github.com/spf13/cobra"
)

var StateCmd = &cobra.Command{
	Use:   "state",
	Short: "Show current state",
	Run: func(_ *cobra.Command, _ []string) {
		s, err := state.GetState()
		if err != nil {
			fmt.Printf("failed to get state: %v\n", err)
			return
		}

		f, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			fmt.Printf("failed to get json state: %v\n", err)
			return
		}

		fmt.Println(string(f))
	},
}
