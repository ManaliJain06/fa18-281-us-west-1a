/*
	go-burger-order REST API (Version)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/handlers"
)

// MongoDB Config
var mongodb_server = "13.57.246.180"
//var mongodb_server = "10.0.0.117"
//var mongodb_server = "dockerhost"
var mongodb_database = "burger"
var mongodb_collection = "order"
var mongo_user = "mongo_admin"
var mongo_pass = "cmpe281"

// RabbitMQ Config
// var rabbitmq_server = "rabbitmq"
// var rabbitmq_port = "5672"
// var rabbitmq_queue = "gumball"
// var rabbitmq_user = "guest"
// var rabbitmq_pass = "guest"

// NewServer configures and returns a Server.
/*func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}*/
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD",  "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders,allowedMethods , allowedOrigins)(router))
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", burgerOrderStatus(formatter)).Methods("GET")
	mx.HandleFunc("/order/{orderId}", burgerOrderStatus(formatter)).Methods("GET")
	mx.HandleFunc("/order", burgerOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/order/{orderId}", burgerOrderPaid(formatter)).Methods("PUT")
	mx.HandleFunc("/order/{orderId}", burgerItemDelete(formatter)).Methods("DELETE")
	mx.HandleFunc("/order", burgerOrderDelete(formatter)).Methods("DELETE")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func getIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
		return "" 
	}
     defer conn.Close()
	 localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	 return localAddr
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"burger-order API is up! " + getIp()})
	}
}

// API burger Order Handler
func burgerOrderStatus(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err:= session.DB("admin").Login(mongo_user, mongo_pass)
		if err!=nil{
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var uuid string = params["orderId"]
		if uuid == "" {
			var orders_array []BurgerOrder
			err = c.Find(bson.M{}).All(&orders_array)
			fmt.Println("Burger Orders:", orders_array)
			formatter.JSON(w, http.StatusOK, orders_array)
		} else {
			fmt.Println("orderID: ", uuid)
			var result BurgerOrder
			err = c.Find(bson.M{"orderId":uuid}).One(&result)
			if err!=nil {
				formatter.JSON(w, http.StatusNotFound, "Order Not Found")
				return
			}
			_ = json.NewDecoder(req.Body).Decode(&result)
			fmt.Println("Burger Order: ", result)
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}

// API Create New Burger Order

func burgerOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session
		var orderdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&orderdetail)
		session, err := mgo.Dial(mongodb_server)
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)		
		c := session.DB(mongodb_database).C(mongodb_collection)
		var order BurgerOrder
		var newitem Items
		// Find order if it exist
		err = c.Find(bson.M{"orderId" : orderdetail.OrderId}).One(&order)
		newitem.ItemId = orderdetail.ItemId
		newitem.ItemName = orderdetail.ItemName
		newitem.Price = orderdetail.Price		
		newitem.Description = orderdetail.Description
		if err == nil {
			order.Cart = append(order.Cart, newitem)
			order.TotalAmount = (order.TotalAmount + newitem.Price)
			fmt.Println("Orders: ", "Orders found")	
			c.Update(bson.M{"orderId": orderdetail.OrderId}, bson.M{"$set": bson.M{"items": order.Cart, "totalAmount": order.TotalAmount, "ipaddress": getIp()}})
		}else {
				fmt.Println("Orders: ", "Orders not found")	
				order = BurgerOrder{
				OrderId:     orderdetail.OrderId,
				UserId:      orderdetail.UserId,
				OrderStatus: "Placed",
				TotalAmount: newitem.Price,
				Cart: []Items{
					newitem,
				},
				IpAddress:	  getIp(), 
			}
			_ = json.NewDecoder(req.Body).Decode(&order)
			err = c.Insert(order)
			if err != nil {
				formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}
		fmt.Println("Orders: ", orders)
		formatter.JSON(w, http.StatusOK, order)
	}
}

// API Paid Order
	func burgerOrderPaid(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var paymentdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&paymentdetail)
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err:= session.DB("admin").Login(mongo_user, mongo_pass)
		if err!=nil{
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var uuid string = params["orderId"]
		fmt.Println(uuid) 
		var orderpaid BurgerOrder
		err = c.Find(bson.M{"orderId": uuid}).One(&orderpaid)
        if err != nil {
			fmt.Println("Order not found") 
			formatter.JSON(w, http.StatusNotFound, "Order Not Found")
			return
        }
		orderpaid.OrderStatus = "Paid"
		orderpaid.UserId = paymentdetail.UserId
		c.Update(bson.M{"orderId": uuid}, bson.M{"$set": bson.M{"orderStatus" : orderpaid.OrderStatus, "userId" : orderpaid.UserId, "ipaddress" : getIp()}})
        fmt.Println("Order:", uuid, "paid" )
		formatter.JSON(w, http.StatusOK, orderpaid)
	} 
} 

// API Delete Item from Order
func burgerItemDelete(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var orderdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&orderdetail)
		params := mux.Vars(req)
		var uuid string = params["orderId"]
		var result BurgerOrder
		fmt.Println("order ID: ", uuid)
		err = c.Find(bson.M{"orderId":uuid}).One(&result)
		if err!=nil{
			fmt.Println("order not found")
			formatter.JSON(w, http.StatusNotFound, "Order Not Found")
			return
		}
		for i := 0; i < len(result.Cart); i++ {
			if result.Cart[i].ItemId == orderdetail.ItemId {
				result.TotalAmount = result.TotalAmount - result.Cart[i].Price
				result.Cart = append(result.Cart[0:i],result.Cart[i+1:]...)
				break
			}
		}
		c.Update(bson.M{"orderId": uuid}, bson.M{"$set": bson.M{"items": result.Cart, "totalAmount": result.TotalAmount, "ipaddress": getIp()}})
		fmt.Println("Delete Item: ", orderdetail.ItemId, "from order", uuid)
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Delete Burger Order
func burgerOrderDelete(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var orderdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&orderdetail)
		fmt.Println("order ID: ", orderdetail.OrderId)
		err = c.Remove(bson.M{"orderId": orderdetail.OrderId})
		if err!=nil{
			fmt.Println("order not found")
			formatter.JSON(w, http.StatusNotFound, "Order Not Found")
			return
		}
		formatter.JSON(w, http.StatusOK, "Order: " + orderdetail.OrderId + " Deleted")
	} 
} 

