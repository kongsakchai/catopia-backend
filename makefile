start:
	docker compose up -d

stop:
	docker compose down

ed:
	nano docker-compose.yml

build-image:
	docker build -t backend .

rm:
	docker rm backend -f

rmi:
	docker rmi backend

exec:
	docker exec -it backend bash

build:
	go build -o main main.go

dev:
	go run main.go

migrate:
	go run ./cmd/migrate.go

swagger:
	swag init