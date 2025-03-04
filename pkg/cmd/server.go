package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/adapter/handler"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := newConfig()
	if err != nil {
		return err
	}

	// di
	server, err := initServer(config)
	if err != nil {
		return err
	}
	// チャネルを作成して OS シグナルを受け取る
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// 別のゴルーチンでサーバーを起動
	go func() {
		slog.Info("Server is running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
		}
	}()

	// OS シグナル (Ctrl+C, SIGTERM) を待つ
	<-stop
	slog.Info("\nShutting down server...")

	// サーバーをシャットダウン
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server gracefully stopped")

	return nil
}

type server struct {
	*http.Server
}

func initServer(cfg *Config) (*server, error) {
	s, err := api.NewServer(handler.NewHandler())
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: s,
	}

	return &server{httpServer}, nil

}
