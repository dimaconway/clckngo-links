
build_output = ./clckngo-links

build: generate-openapi
	go build -o $(build_output) ./cmd/clckngo-links/main.go

run:
	export $$(grep -v '^#' .env | xargs -0) && $(build_output)

build-and-run: build
	$(MAKE) run

generate-openapi:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./api/oapi-codegen_config_types.yml ./api/openapi.yml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./api/oapi-codegen_config_server.yml ./api/openapi.yml
