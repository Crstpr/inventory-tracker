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

.PHONY: run migrate-up migrate-down migrate-status sqlc seed-dev

seed-dev:
	psql "$(ADMIN_DATABASE_URL)" -f ./db/seeds/dev.sql