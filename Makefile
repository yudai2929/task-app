run_server:
	go run cmd/server/main.go server

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

reset:
	@docker compose down
	@docker volume rm task-app_postgres_data

psql:
	docker exec -it postgres_db psql -U postgres -d postgres

test:
	go test ./