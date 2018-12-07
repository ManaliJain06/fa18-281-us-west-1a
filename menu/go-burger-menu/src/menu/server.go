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
	"net"
	"strings"
	"os"
)

// MongoDB Config
//var database_server = "18.144.8.184:27017"
//var database = "burger"
//var collection = "menu"
var database_server = os.Getenv("DatabaseServer")
var database = os.Getenv("Database")
var collection = os.Getenv("Collection")




// MenuServer configures and returns a MenuServer instance.
func MenuServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD","DELETE","OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders,allowedMethods , allowedOrigins)(router))
	return n
}

// Menu Service API Routes
func initRoutes(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/menu/ping", pingHandler(formatter)).Methods("GET")
	router.HandleFunc("/menu", createMenuItemHandler(formatter)).Methods("POST")
	router.HandleFunc("/menu/{restaurantId}", findRestaurantMenu(formatter)).Methods("GET")
	router.HandleFunc("/menu", updateMenuItemHandler(formatter)).Methods("PUT")
	router.HandleFunc("/menu", deleteMenuItemHandler(formatter)).Methods("DELETE")

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
        message := "Burger Menu API Server Working on machine: " + getSystemIp()
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
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
       mongo_collection := session.DB(database).C(collection)
        
       	var menu Menu;
        err = mongo_collection.Find(bson.M{"restaurantid" : reqPayload.RestaurantId}).One(&menu)
        if err != nil {
                fmt.Println("error: ", err)

             	menu.RestaurantId = reqPayload.RestaurantId
             	//menu.RestaurantName = reqPayload.RestaurantName
             	menu.Items = append(menu.Items, reqPayload.Item)
             	
            error := mongo_collection.Insert(menu)
            fmt.Println("error: ", error)
            if error != nil {
                formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
                return
            }
        	
        }else{
        	menu.Items = append(menu.Items, reqPayload.Item)
        	error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": menu.Items}})       	
        	if error != nil {
        		fmt.Println("error: ", error)
                formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
                return
        	}   
        }
		formatter.JSON(response, http.StatusOK, menu)
	}
}


// API to find an item in the menu
func findRestaurantMenu(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		var restaurantId string = params["restaurantId"]
		fmt.Println( "restaurant ID: ", restaurantId )
		session, err := mgo.Dial(database_server)
        if err != nil {
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
        //session.SetMode(mgo.Monotonic, true) need to check
        mongo_collection := session.DB(database).C(collection)
        var result bson.M
        err = mongo_collection.Find(bson.M{"restaurantid" : restaurantId}).One(&result)
        if err != nil {
            formatter.JSON(response, http.StatusNotFound, "Menu not found !!!")
            return
        }
        fmt.Println("Result: ", result)
		formatter.JSON(response, http.StatusOK, result)
	}
}


// API to update an items in the menu
func updateMenuItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var reqPayload restaurantReqBody
		_ = json.NewDecoder(request.Body).Decode(&reqPayload)		
    	fmt.Println("Menu ItemPayload ", reqPayload.Item)
    	session, err := mgo.Dial(database_server)
        if err != nil {
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
       mongo_collection := session.DB(database).C(collection)
        
       	var menu Menu;
        err = mongo_collection.Find(bson.M{"restaurantid" : reqPayload.RestaurantId}).One(&menu)
        if err != nil {
            fmt.Println("error: ", err)
            formatter.JSON(response, http.StatusNotFound, "Restaurant not found")
        	return
        }else{
        	for i := 0; i < len(menu.Items); i++ {
				if menu.Items[i].Id == reqPayload.Item.Id {
					menu.Items[i].Name = reqPayload.Item.Name
					menu.Items[i].Price = reqPayload.Item.Price
					menu.Items[i].Description = reqPayload.Item.Description
					menu.Items[i].Calories = reqPayload.Item.Calories
					break
				}
			}
        	error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": menu.Items}})  
        	if error != nil {
        		fmt.Println("error: ", error)
                formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
                return
        	}     	
        }
        
		formatter.JSON(response, http.StatusOK, menu)
	}
}

// API to delete an items in the menu
func deleteMenuItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var reqPayload deleteReqBody
		_ = json.NewDecoder(request.Body).Decode(&reqPayload)		
    	fmt.Println("Menu ItemPayload ", reqPayload)
    	session, err := mgo.Dial(database_server)
        if err != nil {
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
        mongo_collection := session.DB(database).C(collection)
        
       	var menu Menu; 
        err = mongo_collection.Find(bson.M{"restaurantid" : reqPayload.RestaurantId}).One(&menu)
        if err != nil {
            fmt.Println("error: ", err)
            formatter.JSON(response, http.StatusNotFound, "Restaurant not found")
        	return
        }else{
        	for i := 0; i < len(menu.Items); i++ {
				if menu.Items[i].Id == reqPayload.ItemId {
					menu.Items = append(menu.Items[0:i],menu.Items[i+1:]...)
					break
				}
			}
        	error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": menu.Items}})  
        	if error != nil {
        		fmt.Println("error: ", error)
                formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
                return
        	}     	
        }
		formatter.JSON(response, http.StatusOK, menu)
	}
}





