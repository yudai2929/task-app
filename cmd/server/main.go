package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yudai2929/task-app/pkg/cmd"
)

func main() {
	c := &cobra.Command{} //nolint:exhaustruct
	cmd.RegisterCommand(c)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
