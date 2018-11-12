/*
All the REST API calls for location module of the counter burger restaurant location will be handled here
Referenced from- https://github.com/paulnguyen/cmpe281/blob/master/golabs/godata/go-gumball-mongo/src/gumball/server.go
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

// Configuration parameters for MongoDB databse
var mongodb_server = "localhost"
var mongodb_database = "burger"
var mongodb_collection = "restaurantLocation"

/*
Reference for Server configuration taken from - https://github.com/paulnguyen/cmpe281/blob/master/golabs/godata/go-gumball-mongo/src/gumball/server.go
*/
func NewServerConfiguration() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// Router for API
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/restaurant", addRestaurantHandler(formatter)).Methods("POST")
}

// Handler for API Ping
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

/*
Handler method for adding restaurant
*/
func addRestaurantHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		uuidForRestaurant,_ := uuid.NewV4()

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var res restaurant
		_=json.NewDecoder(req.Body).Decode(&res)
		
		res.RestaurantId = uuidForRestaurant.String()
		fmt.Println("Restuanats: ", res)
		err = c.Insert(res)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, res)
	}
}
