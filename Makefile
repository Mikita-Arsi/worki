stop:
	docker-compose stop
down:
	docker-compose down
build:
	docker-compose --env-file .env build
up:
	docker-compose up
run: build up

rerun: stop down run

set-debez-conf: 
	curl -i -X POST \
         -H "Accept:application/json" \
         -H "Content-Type:application/json" \
         127.0.0.1:8083/connectors/ \
         --data "@debezium.json"

all: rerun set-debez-co