go-get:
	rm -rf src/github.com
	go get -v github.com/codegangsta/negroni
	go get -v github.com/streadway/amqp
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v github.com/gorilla/handlers
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson


run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt menu

install:
	go install menu

build:
	go build menu

start:
	./menu

docker-build:
	docker build -t menu .
	docker images

docker-run:
	docker run --name menu -td -p 8000:8000 menu
	docker ps

menu-tag:
	docker tag menu arijitmandal/test:menu-aws
	docker images

push:
	docker push arijitmandal/test

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it menu bash

docker-clean:
	docker stop menu
	docker rm menu
	docker rmi menu

docker-ip:
	docker-machine ip
