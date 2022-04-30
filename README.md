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
export export YOUR_VOLUME_NAME=my_project_mysql

docker volume create --name=$YOUR_VOLUME_NAME
sed /MUST_BE_REPLACED/s//$YOUR_VOLUME_NAME/ docker-compose.yml.sample > docker-compose.yml

cp env.sample .env
docker-compose up
```

## Command
```
docker-compose exec db mysql -u root -p
```
