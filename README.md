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
cp build/package/env.sample build/package/env
export YOUR_VOLUME_NAME=my_project_mysql

docker volume create --name=$YOUR_VOLUME_NAME
sed s/MUST_BE_REPLACED/$YOUR_VOLUME_NAME/ build/package/docker-compose.yml.sample > build/package/docker-compose.yml
unset YOUR_VOLUME_NAME

docker-compose -f build/package/docker-compose.yml up
```

## Command
```
docker-compose exec db mysql -u root -p
```
