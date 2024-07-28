dev-api:
	go run cmd/api/main.go

dev-consumer:
	go run cmd/consumer/main.go

dev-cli:
	go run cmd/cli/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down