## Cloud Disk Demo
Lightweight cloud disk demo, based on go-zero, MySQL, Redis and Gorm and using AWS S3

```
# use docker to create mysql
docker run --name mysql-db -p 3306:3306  -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=zero-todo -e MYSQL_USER=user -e MYSQL_PASSWORD=password -d mysql:latest

# use docker to create redis
docker run --name redis -d -p 6379:6379 redis:latest

# create api service
goctl api new core

# start service
go run core.go -f etc/core-api.yaml

# use .api to generate code
goctl api go -api core.api -dir . -style go_zero
```
## Main Features
- User module
  - Password login
  - Authorization
  - User details
  - Register by receiving an email code
- Repository module
  - Connect to AWS S3 bucket
  - File upload
  - File multipart upload
  - File list
  - File associated storage
  - Folder creation
