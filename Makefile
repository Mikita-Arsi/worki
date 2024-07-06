kill:
	docker-compose kill
down:
	docker-compose down
build:
	docker-compose --env-file .env build
up:
	docker-compose up
run: build up

rerun: kill down run