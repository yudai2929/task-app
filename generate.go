package taskapp

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target doc/gen --clean doc/api.yml
//go:generate go run github.com/xo/xo@latest schema postgres://postgres:password@localhost:5432/postgres?sslmode=disable -o database/gen
