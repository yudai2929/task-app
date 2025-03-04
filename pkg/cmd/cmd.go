package cmd

import (
	"github.com/spf13/cobra"
)

// RegisterCommand registers the command to the registry
func RegisterCommand(registry *cobra.Command) {
	registry.AddCommand(
		newServerApp().Command,
	)
}
