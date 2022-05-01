#!/usr/bin/env bash

rm -rf internal/pkg/openapi

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i local/api/swagger.yml -g go-server -o /local/internal/pkg/openapi --package-name openapi

rm -f internal/pkg/openapi/go.mod
rm -f internal/pkg/openapi/go.sum
rm -rf internal/pkg/openapi/docs
rm -rf internal/pkg/openapi/api
rm -f internal/pkg/openapi/git_push.sh
rm -f internal/pkg/openapi/.gitignore
rm -f internal/pkg/openapi/.travis.yml
rm -f internal/pkg/openapi/.openapi-generator/FILES
rm -f internal/pkg/openapi/Dockerfile
rm -f internal/pkg/openapi/main.go

docker-compose exec api go fmt ./...
