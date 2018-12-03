/*
	Burger Menu Item API
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"github.com/gorilla/handlers"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var database_server = "18.144.8.184:27017"
var database = "burger"
var collection = "menu"



// MenuServer configures and returns a MenuServer instance.
func MenuServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders,allowedMethods , allowedOrigins)(router))
	return n
}

// Menu Service API Routes
func initRoutes(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/menu/ping", pingHandler(formatter)).Methods("GET")
	router.HandleFunc("/menu/item", createMenuItemHandler(formatter)).Methods("POST")
	router.HandleFunc("/menu/item/{id}", findItemHandler(formatter)).Methods("GET")
	router.HandleFunc("/menu/item", updateItemsHandler(formatter)).Methods("PUT")

}

// Error Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// Menu Serivce Health Check API 
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Burger Menu API Server Working !!!"})
	}
}

// API to create a new item in the menu
/*func createMenuItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		uuid,_ := uuid.NewV4()
		var item menuItem
		_ = json.NewDecoder(request.Body).Decode(&item)		
    	fmt.Println("Menu Item: ", item.Name)
    	fmt.Println("Menu Item Id: ", uuid)
    	fmt.Println("Menu Item : ", item)
		session, err := mgo.Dial(database_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        //session.SetMode(mgo.Monotonic, true) need to check
        item.Id = uuid.String()
        mongo_collection := session.DB(database).C(collection)
        error := mongo_collection.Insert(item)
        if error != nil {
                panic(error)
        }
        fmt.Println("Menu mongo_collection: ", mongo_collection)
		formatter.JSON(response, http.StatusOK, item)
	}
}*/

// API to create a new item in the menu
func createMenuItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var reqPayload restaurantReqBody
		_ = json.NewDecoder(request.Body).Decode(&reqPayload)		
    	fmt.Println("Menu ItemPayload ", reqPayload.Item)
    	uuid,_ := uuid.NewV4()
    	reqPayload.Item.Id = uuid.String()
    	session, err := mgo.Dial(database_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
       mongo_collection := session.DB(database).C(collection)
        //var newMenu bson.M

       	var menu Menu;
        err = mongo_collection.Find(bson.M{"restaurantid" : reqPayload.RestaurantId}).One(&menu)
        if err != nil {
                fmt.Println("error: ", err)

             	menu.RestaurantId = reqPayload.RestaurantId
             	menu.RestaurantName = reqPayload.RestaurantName
             	menu.Items = append(menu.Items, reqPayload.Item)
             	
            error := mongo_collection.Insert(menu)
            fmt.Println("error: ", error)
        	
        }else{
        	menu.Items = append(menu.Items, reqPayload.Item)
        	error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": append(menu.Items, reqPayload.Item)}})       	
        	fmt.Println("error: ", error)
        }
        
        
		formatter.JSON(response, http.StatusOK, menu)
	}
}


// API to find an item in the menu
func findItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		var uuid string = params["id"]
		fmt.Println( "Item ID: ", uuid )
		session, err := mgo.Dial(database_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        //session.SetMode(mgo.Monotonic, true) need to check
        mongo_collection := session.DB(database).C(collection)
        var result bson.M
        err = mongo_collection.Find(bson.M{"id" : uuid}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("result: ", result)
		formatter.JSON(response, http.StatusOK, result)
	}
}

// API to update an items in the menu
func updateItemsHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		uuid,_ := uuid.NewV4()
		var item menuItem
		_ = json.NewDecoder(request.Body).Decode(&item)		
    	fmt.Println("Menu Item: ", item.Name)
    	fmt.Println("Menu Item Id: ", uuid)
    	fmt.Println("Menu Item : ", item)
		session, err := mgo.Dial(database_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        //session.SetMode(mgo.Monotonic, true) need to check
        item.Id = uuid.String()
        mongo_collection := session.DB(database).C(collection)
        error := mongo_collection.Insert(item)
        if error != nil {
                panic(error)
        }
        fmt.Println("Menu mongo_collection: ", mongo_collection)
		formatter.JSON(response, http.StatusOK, item)
	}
}





