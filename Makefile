run	:
	go run cmd/main.go
up	:
	docker-compose -f deployments/docker-compose.yaml up -d
down	:
	docker-compose -f deployments/docker-compose.yaml down
volume	:
	docker volume create --name=crud-api
