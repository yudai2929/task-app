run_server:
	go run cmd/server/main.go

set_up:
	go mod tidy

install:
	@go install github.com/ogen-go/ogen/cmd/ogen@latest
	@go install github.com/xo/xo@latest

gen:
	go generate ./...

up:
	docker compose up

down:
	docker compose down

psql:
	docker exec -it postgres_db psql -U user -d task_db