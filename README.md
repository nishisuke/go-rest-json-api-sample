# go-rest-json-api-sample
## GitHub settings
### Protection
https://github.com/nishisuke/go-rest-json-api-sample/settings/branches

- Require a pull request before merging
  - Require approvals
  - Dismiss stale pull request approvals when new commits are pushed
- Include administrators
### Workflow permissions
https://github.com/nishisuke/go-rest-json-api-sample/settings/actions

- Read and write permissions


## TODOs
- x swagger.yml
- x golangci-lint
- x test

## Versions
- Go 1.18
- Air any

## Rules
### Directories
https://github.com/golang-standards/project-layout

## Setup
```
# Define your volume name
export YOUR_PROJECT_NAME_SNAKE_CASE=my_project

docker volume create --name=$YOUR_PROJECT_NAME_SNAKE_CASE
sed s/MUST_BE_REPLACED/$YOUR_PROJECT_NAME_SNAKE_CASE/ build/package/docker-compose.yml.sample > docker-compose.yml
sed s/MUST_BE_REPLACED/${YOUR_PROJECT_NAME_SNAKE_CASE}_db/ build/package/env.sample > .env

unset YOUR_PROJECT_NAME_SNAKE_CASE

docker-compose up
```

## Command
```
docker-compose exec db mysql -u root -p
docker-compose exec api golangci-lint run
docker-compose exec api swag i -o api --ot go,yaml -g cmd/rest-json-api/main.go
docker run --rm -v `pwd`:/local openapitools/openapi-generator-cli generate -i local/api/swagger.yaml -g go -o /local/pkg/openapicli --package-name openapicli
docker run --rm -v `pwd`:/local openapitools/openapi-generator-cli generate -i local/api/swagger.yaml -g typescript-axios -o /local/web/src
```
# go-rest-json-api-sample

dc exec api gocyclo -avg  -ignore "pkg/openapicli/" .
// 3 はsafe
dc exec api gocyclo -over 3 -ignore "pkg/openapicli/" .       

