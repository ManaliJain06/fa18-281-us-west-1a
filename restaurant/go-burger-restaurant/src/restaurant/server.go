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
	"github.com/gorilla/handlers"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
	"strings"
)

// Configuration parameters for MongoDB databse
var mongodb_server = "54.67.41.59:27017"
var mongodb_database = "burger"
var mongodb_collection = "restaurant"
var mongo_user = "mongo-admin"
var mongo_pass = "cmpe281"
var adminDatabase = "admin"
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
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})

    n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(mx))
	return n
}

// Router for API
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/restaurant/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/restaurant", addRestaurantHandler(formatter)).Methods("POST")
	mx.HandleFunc("/restaurant", getAllRestaurantHandler(formatter)).Methods("GET")
	mx.HandleFunc("/restaurant/{restaurantId}", getRestaurantByIDHandler(formatter)).Methods("GET")
	mx.HandleFunc("/restaurant/{restaurantId}", deleteRestaurantHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/restaurant/zipcode/{zipcode}", getRestaurantHandler(formatter)).Methods("GET")
	// mx.handleFunc("/restaurant/{restaurantId}", updateRestaurantHandler(formatter)).Methods("PUT")
}

// Handler for API Ping
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Burger Restaurant API Server Working on machine: " + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

func getSystemIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
		return "" 
	}
    defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	address := strings.Split(localAddr, ":")
    fmt.Println("address: ", address[0])
    return address[0]
}

/*
Handler method for adding restaurant
*/
func addRestaurantHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		uuidForRestaurant,_ := uuid.NewV4()
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		var res restaurant
		_=json.NewDecoder(req.Body).Decode(&res)
		
		res.RestaurantId = uuidForRestaurant.String()
		fmt.Println("Restaurants: ", res)
		err = collection.Insert(res)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Error occureed. Cannot add restaurant")
			return
		} 
		formatter.JSON(w, http.StatusOK, res)
	}
}

/*
Handler method for getting all restaurants based on a ziplocation
*/
func getRestaurantHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)
		var zipcode string = params["zipcode"]
		fmt.Println(zipcode);
		
		var res  []restaurant
        err = collection.Find(bson.M{"zipcode" : zipcode}).All(&res)
		
		if res == nil || len(res) <= 0 || err != nil{
			formatter.JSON(w, http.StatusNotFound, "Cannot find any restaurants for that zipcode") 
		} else {
			fmt.Println("Result: ", res)
			formatter.JSON(w, http.StatusOK, res)
		}
	}
}

/*
Handler to get all restaurant
*/
func getAllRestaurantHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		// params := mux.Vars(req)
		// var zipcode string = params["zipcode"]
		// fmt.Println(zipcode);
		
		var res  []restaurant
        err = collection.Find(bson.M{}).All(&res)
		
		if res == nil || len(res) <= 0 || err != nil{
			formatter.JSON(w, http.StatusNotFound, "No restaurants found") 
		} else {
			fmt.Println("Result: ", res)
			formatter.JSON(w, http.StatusOK, res)
		}
	}
}

/*
Handler method for getting restaurant based on a Id
*/
func getRestaurantByIDHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			log.Fatal(err);
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)
		var restaurantId string = params["restaurantId"]
		fmt.Println("restaurant id is : ", restaurantId)

		var res restaurant
        err = collection.Find(bson.M{"restaurantid" : restaurantId}).One(&res)
		
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Cannot find restaurant")
			return
		} else {	
			fmt.Println("Result: ", res)
			// res := json.NewEncoder(w).Encode(res)
			formatter.JSON(w, http.StatusOK, res)
		}
	}
}
func deleteRestaurantHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		
		params := mux.Vars(req)

		var result restaurant
		err = c.Find(bson.M{"restaurantid": params["restaurantId"]}).One(&result)
		if err == nil {
			c.Remove(bson.M{"restaurantid": params["restaurantId"]})
			formatter.JSON(w, http.StatusOK, result)
		} else {
			formatter.JSON(w, http.StatusNotFound, "Restaurant not found for delete")
		}	
	}
}

// func updateRestaurantHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
		
// 		session, err := mgo.Dial(mongodb_server)
// 		if err != nil {
// 			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
// 			return
// 		}
// 		defer session.Close()

// 		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
// 			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
// 			return
// 		  }
// 		session.SetMode(mgo.Monotonic, true)
// 		c := session.DB(mongodb_database).C(mongodb_collection)
		
// 		params := mux.Vars(req)

// 		var newRes restaurant
// 		_=json.NewDecoder(req.Body).Decode(&newRes)

// 		var newRestaurant bson.M
// 		err = c.Find(bson.M{"restaurantId":res.restaurantId}).One(&newRestaurant)
// 		if err != nil {
// 			formatter.JSON(w, http.StatusNotFound, "Restaurant not found for update")
// 			return 
// 		} else {
// 			query := bson.M{"id":params["id"]}
// 			updatorQuery := bson.M{"$set": bson.M{"restaurantName": res.restaurantName
// 						"zipcode": res.zipcode
// 						"phone": res.phone
// 						"addressLine1": res.addressLine1
// 						"addressLine2": res.addressLine2
// 						"city": res.city
// 						"state": res.state
// 						"country": res.country
// 						"hours": res.hours
// 						"acceptedCards": res.acceptedCards
// 						"distance": res.distance
// 						"email": res.email}}
//     		err = c.Update(query, updatorQuery)

//     		if err != nil {
// 				formatter.JSON(w, http.StatusInternalServerError, "Error Occured while Updating")
// 				return
// 			} else {
// 				formatter.JSON(w, http.StatusOK, newRestaurant)
// 			}
// 		}
// 	}
// }