-include .env
export

DB_URL ?= postgres://postgres:your_password@localhost:5432/my-api?sslmode=disable

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir database/migrations -seq $(name)

generate-mock:
	mockery --all
