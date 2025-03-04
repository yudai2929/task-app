package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type App struct {
	*cobra.Command
}

func newServerApp() *App {
	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "server",
		Short: "Run the server",
	}
	app := &App{Command: cmd}
	app.RunE = func(_ *cobra.Command, _ []string) error {
		return app.run()
	}
	return app
}

func (s *App) run() error {
	fmt.Println("Hello, World!")
	return nil
}
