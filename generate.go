package taskapp

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target doc/gen --clean doc/api.yml
//go:generate go run github.com/xo/xo@latest schema postgres://user:password@localhost:5432/task_db?sslmode=disable -o database/gen
