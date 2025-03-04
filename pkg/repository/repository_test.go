package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var db *sql.DB

func TestMain(m *testing.M) {
	fmt.Println("🔹 TestMain started")

	// Docker プールの作成
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// PostgreSQL のコンテナ起動
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env: []string{
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=postgres",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// コンテナのポートを取得
	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseURL := fmt.Sprintf("postgres://postgres:postgres@%s/postgres?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseURL)

	// DB 接続をリトライ
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseURL)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	fmt.Println("✅ Database connected")

	// マイグレーション適用
	if err := applyMigrations(db, "../../database/migration"); err != nil {
		log.Fatalf("Could not apply migrations: %s", err)
	}

	// defer で DB をクローズ
	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()

	// テスト実行
	code := m.Run()

	// コンテナ削除
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	fmt.Println("🛑 TestMain finished")

	os.Exit(code)
}

func applyMigrations(db *sql.DB, migrationDir string) error {
	fmt.Println("🚀 Applying migrations...")

	files, err := filepath.Glob(filepath.Join(migrationDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("could not read migration files: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no migration files found in %s", migrationDir)
	}

	for _, file := range files {
		sqlContent, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("could not read SQL file %s: %w", file, err)
		}

		fmt.Printf("📝 Executing migration: %s\n", file)
		if _, err := db.Exec(string(sqlContent)); err != nil {
			return fmt.Errorf("could not execute migration %s: %w", file, err)
		}
	}

	fmt.Println("✅ Migrations applied successfully")
	return nil
}
