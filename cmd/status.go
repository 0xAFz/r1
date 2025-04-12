package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/0xAFz/r1/internal/state"
	"github.com/spf13/cobra"
)

var StateCmd = &cobra.Command{
	Use:   "state",
	Short: "Shows the attributes of resources in the R1 state.",
	Run: func(_ *cobra.Command, _ []string) {
		s, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		f, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			fmt.Printf("marshaling current state: %v\n", err)
			return
		}

		fmt.Println(string(f))
	},
}
