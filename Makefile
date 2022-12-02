.PHONY: test
run	:
	go run cmd/main.go

test	:
	go test -cover ./...

up	:
	docker-compose -f deployments/docker-compose.yaml up -d

down	:
	docker-compose -f deployments/docker-compose.yaml down

volume	:
	docker volume create --name=crud-api

lint	:
	golangci-lint run --timeout 5m0s

swag:
	swag init --parseDependency --parseInternal -g advertise.go -d ./internal/server,./internal/model -o api -ot yaml,go & \
	swag fmt

docker	:
	docker build -t crud-api .
