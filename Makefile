include .env
export $(shell sed 's/=.*//' .env)

build_output = ./clckngo-links

build:
	go build -o $(build_output) ./cmd/clckngo-links/main.go

run:
	export $$(grep -v '^#' .env | xargs -0) && $(build_output)

build-and-run: build
	$(MAKE) run

watch:
	go run github.com/cosmtrek/air

mod:
	go mod tidy
	go mod vendor

openapi:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./api/oapi-codegen_config_types.yml ./api/openapi.yml

local-migrate-up:
	migrate -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${DB_PORT}/${POSTGRES_DB}?sslmode=disable -path db/migrations up 1

local-migrate-down:
	migrate -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${DB_PORT}/${POSTGRES_DB}?sslmode=disable -path db/migrations down 1
