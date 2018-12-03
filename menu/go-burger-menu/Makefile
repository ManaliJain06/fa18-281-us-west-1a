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
