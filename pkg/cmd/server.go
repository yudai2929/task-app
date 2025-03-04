package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/adapter/handler"
	"github.com/yudai2929/task-app/pkg/adapter/middleware"
	"github.com/yudai2929/task-app/pkg/repository"
	"github.com/yudai2929/task-app/pkg/usecase"

	_ "github.com/lib/pq"
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
	server, cleanup, err := initServer(config)
	if err != nil {
		return err
	}
	defer cleanup()

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

func initServer(cfg *Config) (*server, func(), error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	ur := repository.NewUserRepository(db)
	au := usecase.NewAuthUsecase(ur, cfg.JWTSecret, cfg.TokenExpiry)
	s, err := api.NewServer(handler.NewHandler(au), api.WithMiddleware(middleware.AccessLog()))
	if err != nil {
		return nil, nil, err
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // TODO: 許可するドメインを指定
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(s)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: corsHandler,
	}

	return &server{httpServer}, func() {
		db.Close()
	}, nil

}
