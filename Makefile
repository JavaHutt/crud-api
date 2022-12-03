.PHONY: test
run	:
	go run cmd/main.go

test	:
	CGO_ENABLED=1 go test -race -cover -count=1 -coverprofile=./test/.coverprofile ./internal/...

up	:
	docker-compose -f deployments/docker-compose.yaml up -d

down	:
	docker-compose -f deployments/docker-compose.yaml down

volume	:
	docker volume create --name=crud-api

fake	:
	curl --location --request GET 'http://localhost:3000/api/v1/faker?num=50'

lint	:
	golangci-lint run --timeout 5m0s

swag	:
	swag init --parseDependency --parseInternal -g query.go -d ./internal/server,./internal/model -o api -ot yaml,go & \
	swag fmt

docker	:
	docker build -t crud-api .
