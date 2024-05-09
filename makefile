start:
	docker compose up -d

stop:
	docker compose down

edit:
	nano docker-compose.yml

build:
	docker build -t backend .

rm:
	docker rm backend -f

rmi:
	docker rmi backend

exec:
	docker exec -it backend bash