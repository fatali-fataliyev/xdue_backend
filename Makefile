# Load .env
ifneq ("$(wildcard .env)", "")
	include .env
	export
endif


.PHONY: migrate-up migrate-down migrate-force migrate-version

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path ./migrations -database "$(DB_URL)" force $V

migrate-v:
	migrate -path ./migrations -database "$(DB_URL)" version