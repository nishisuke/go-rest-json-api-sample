# go-rest-json-api-sample
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
sed s/MUST_BE_REPLACED/$YOUR_PROJECT_NAME_SNAKE_CASE/ build/package/docker-compose.yml.sample > build/package/docker-compose.yml
sed s/MUST_BE_REPLACED/${YOUR_PROJECT_NAME_SNAKE_CASE}_db/ build/package/env.sample > build/package/.env

unset YOUR_PROJECT_NAME_SNAKE_CASE

docker-compose -f build/package/docker-compose.yml up
```

## Command
```
docker-compose exec db mysql -u root -p
```
