.PHONY: run migrate-up migrate-down migrate-status

run:
	go run ./cmd/api

migrate-up:
	goose -dir ./db/migrations postgres "$(ADMIN_DATABASE_URL)" up

migrate-down:
	goose -dir ./db/migrations postgres "$(ADMIN_DATABASE_URL)" down

migrate-status:
	goose -dir ./db/migrations postgres "$(ADMIN_DATABASE_URL)" status

sqlc:
	sqlc generate