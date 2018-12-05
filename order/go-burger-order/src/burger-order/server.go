/*
	go-burger-order REST API (Version)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/handlers"
)

// MongoDB Config
var mongodb_server = "13.57.246.180"
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
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
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
	// mx.HandleFunc("/order/{OrderId}", burgerOrderUpdate(formatter)).Methods("PUT")
	mx.HandleFunc("/order/{orderId}", burgerOrderDelete(formatter)).Methods("DELETE")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"burger-order API is up!"})
	}
}

// API burger Order Handler
func burgerOrderStatus(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			panic(err)
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var uuid string = params["orderId"]
		if uuid == "" {
			var orders_array []BurgerOrder
			err = c.Find(bson.M{}).All(&orders_array)
			if err != nil {
				panic(err)
			}
			fmt.Println("Burger Orders:", orders_array)
			formatter.JSON(w, http.StatusOK, orders_array)
		} else {
			fmt.Println("orderID: ", uuid)
			var result BurgerOrder
			err = c.Find(bson.M{"orderId":uuid}).One(&result)
			if err != nil {
				panic(err)
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
			panic(err)
		  }
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)		
		c := session.DB(mongodb_database).C(mongodb_collection)
		uuid2, _ := uuid.NewV4()
		var order BurgerOrder
		var newitem Items
		err = c.Find(bson.M{"orderId" : orderdetail.OrderId}).One(&order)
		newitem.ItemId = orderdetail.ItemId
		newitem.ItemName = orderdetail.ItemName
		newitem.Price = orderdetail.Price		
		newitem.Description = orderdetail.Description
		if err == nil {
			order.Cart = append(order.Cart, newitem)
			order.TotalAmount = (order.TotalAmount + newitem.Price)
			fmt.Println("Orders: ", "Orders found")	
			c.Update(bson.M{"orderId": orderdetail.OrderId}, bson.M{"$set": bson.M{"items": order.Cart}})
			c.Update(bson.M{"orderId": orderdetail.OrderId}, bson.M{"$set": bson.M{"totalAmount": order.TotalAmount}})
		}else {
				fmt.Println("Orders: ", "Orders not found")	
				order = BurgerOrder{
				OrderId:     orderdetail.OrderId,
				UserId:      uuid2.String(),
				OrderStatus: "Order Placed",
				TotalAmount: newitem.Price,
				Cart: []Items{
					newitem,
				},
			}
			_ = json.NewDecoder(req.Body).Decode(&order)
			err = c.Insert(order)
			if err != nil {
				log.Fatal(err)
			}
		}
		/*if orders == nil {
			orders = make(map[string]BurgerOrder)
		}*/
		fmt.Println("Orders: ", orders)
		formatter.JSON(w, http.StatusOK, order)
	}
}
// API Delete Burger Order
func burgerOrderDelete(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			panic(err)
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
			panic(err)
		}
		for i := 0; i < len(result.Cart); i++ {
			if result.Cart[i].ItemId == orderdetail.ItemId {
				result.TotalAmount = result.TotalAmount - result.Cart[i].Price
				result.Cart = append(result.Cart[0:i],result.Cart[i+1:]...)
				break
			}
		}
		c.Update(bson.M{"orderId": uuid}, bson.M{"$set": bson.M{"items": result.Cart}})
		c.Update(bson.M{"orderId": uuid}, bson.M{"$set": bson.M{"totalAmount": result.TotalAmount}})
		//err = c.Remove(bson.M{"orderId": uuid})
		fmt.Println("Delete Item: ", orderdetail.ItemId, "from order", uuid)
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Update Burger Order
/*	func burgerOrderUpdate(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
    	var m BurgerOrder
    	_ = json.NewDecoder(req.Body).Decode(&m)
    	// fmt.Println("Update Gumball Inventory To: ", m.CountGumballs)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			panic(err)
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var uuid string = params["OrderId"]
        query := bson.M{"OrderId" : uuid}
        change := bson.M{"$set": bson.M{ "CountGumballs" : m.CountGumballs}}
        // err = c.Update(query, change)
        if err != nil {
                log.Panic(err)
        }
       	var result bson.M
        err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Gumball Machine:", result )
		formatter.JSON(w, http.StatusOK, result)
	} 
} */

// API Process Orders
/* func gumballProcessOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

       	// Get Gumball Inventory
        var result bson.M
        err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

 		var count int = result["CountGumballs"].(int)
        fmt.Println("Current Inventory:", count )

		// Process Order IDs from Queue
		var order_ids []string = queue_receive()
		for i := 0; i < len(order_ids); i++ {
			var order_id = order_ids[i]
			fmt.Println("Order ID:", order_id)
			var ord = orders[order_id]
			ord.OrderStatus = "Order Processed"
			orders[order_id] = ord
			count -= 1
		}
		fmt.Println( "Orders: ", orders , "New Inventory: ", count)

		// Update Gumball Inventory
		query := bson.M{"SerialNumber" : "1234998871109"}
        change := bson.M{"$set": bson.M{ "CountGumballs" : count}}
        err = c.Update(query, change)
        if err != nil {
                log.Fatal(err)
        }

		// Return Order Status
		formatter.JSON(w, http.StatusOK, orders)
	}
} */

// Send Order to Queue for Processing
/* func queue_send(message string) {
	conn, err := amqp.Dial("amqp://"+rabbitmq_user+":"+rabbitmq_pass+"@"+rabbitmq_server+":"+rabbitmq_port+"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
*/
// Receive Order from Queue to Process
/* func queue_receive() []string {
	conn, err := amqp.Dial("amqp://"+rabbitmq_user+":"+rabbitmq_pass+"@"+rabbitmq_server+":"+rabbitmq_port+"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"orders",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	order_ids := make(chan string)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			order_ids <- string(d.Body)
		}
		close(order_ids)
	}()

	err = ch.Cancel("orders", false)
	if err != nil {
	    log.Fatalf("basic.cancel: %v", err)
	}

	var order_ids_array []string
	for n := range order_ids {
    	order_ids_array = append(order_ids_array, n)
    }

    return order_ids_array
} */

/*

	-- Gumball MongoDB Create Database

		Database Name: cmpe281
		Collection Name: gumball

  	-- Gumball MongoDB Collection (Create Document) --

    db.order.insert(
	    {
	      OrderId: "testorder1",
	      UserId: "testUser1",
	      OrderStatus: 'Test',
	      TotalAmount: NumberInt(10)
	    }
	) ;

    -- Gumball MongoDB Collection - Find Gumball Document --

    db.gumball.find( { Id: 1 } ) ;

    {
        "_id" : ObjectId("54741c01fa0bd1f1cdf71312"),
        "Id" : 1,
        "CountGumballs" : 202,
        "ModelNumber" : "M102988",
        "SerialNumber" : "1234998871109"
    }

    -- Gumball MongoDB Collection - Update Gumball Document --

    db.gumball.update(
        { Dd: 1 },
        { $set : { CountGumballs : NumberInt(10) } },
        { multi : false }
    )

    -- Gumball Delete Documents

    db.gumball.remove({})

*/
