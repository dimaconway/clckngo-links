
generate-openapi:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./api/oapi-codegen_config_types.yml ./api/openapi.yml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./api/oapi-codegen_config_server.yml ./api/openapi.yml
