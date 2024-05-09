start:
	docker compose up -d

stop:
	docker compose down

edit:
	nano docker-compose.yml

run:
	docker run --name backend -v /others:/root/others  --env-file ./.env --network host -i -d backend

build:
	docker build -t backend .

rm:
	docker rm backend -f

rmi:
	docker rmi backend

exec:
	docker exec -it backend bash