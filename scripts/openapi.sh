#!/usr/bin/env bash

rm -rf internal/pkg/openapi
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i local/api/swagger.yml -g go -o /local/internal/pkg/openapi --package-name openapi
rm -f internal/pkg/openapi/go.mod
rm -f internal/pkg/openapi/go.sum
