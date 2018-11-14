/*

 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "localhost:27017"
var mongodb_database = "test"
var mongodb_collection = "payments"

type Payments []Payment

var payments []Payment

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	t := time.Now()
	payments = append(payments, Payment{PaymentID: "1", UserID: "1", OrderID: "11", TotalAmount: 100.50, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})
	payments = append(payments, Payment{PaymentID: "2", UserID: "2", OrderID: "22", TotalAmount: 30.30, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})

	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payments", getAllPayments(formatter)).Methods("GET")
	mx.HandleFunc("/payments/{id}", getPaymentByID(formatter)).Methods("GET")
	mx.HandleFunc("/payments", createPayments(formatter)).Methods("POST")
	mx.HandleFunc("/payments/{id}", deletePayment(formatter)).Methods("DELETE")
	mx.HandleFunc("/payments/{id}", editPayment(formatter)).Methods("PUT")

}

func handleRequest() {

	// myRouter := mux.NewRouter().StrictSlash(true)
	t := time.Now()
	payments = append(payments, Payment{PaymentID: "1", UserID: "1", OrderID: "11", TotalAmount: 100.50, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})
	payments = append(payments, Payment{PaymentID: "2", UserID: "2", OrderID: "22", TotalAmount: 30.30, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})

	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/payments", allPayments).Methods("GET")
	// myRouter.HandleFunc("/payments/{id}", getPaymentById).Methods("GET")
	// myRouter.HandleFunc("/payments", createPayments).Methods("POST")
	// myRouter.HandleFunc("/payments/{id}", deletePayment).Methods("DELETE")

	// log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Payment API version 1.0 is alive!"})
	}
}

// func allPayments(w http.ResponseWriter, r *http.Request) {
// t := time.Now()
// payments := Payments{
// 	Payment{paymentId: "100", userId: "1", orderId: "33", totalAmount: 100.50, status: true, paymentDate: t.Format("2006-01-02 15:04:05")},
// }

// payments := Payments{
// 	Payment{PaymentID: "100", UserID: "1", OrderID: "33", TotalAmount: 100.50, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")},
// }

// 	fmt.Println("Get all payments")

// 	fmt.Println(payments)
// 	json.NewEncoder(w).Encode(payments)
// }

func getAllPayments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result []bson.M
		err = c.Find(nil).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("getAllPayments:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

func getPaymentByID(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result bson.M
		params := mux.Vars(req)
		err = c.Find(bson.M{"PaymentID": params["id"]}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Gumball Machine:", result)
		formatter.JSON(w, http.StatusOK, result)

		// fmt.Println("getPaymentByPaymentId")
		// params := mux.Vars(req)
		// for _, item := range payments {
		// 	if item.PaymentID == params["id"] {
		// 		// json.NewEncoder(w).Encode(item)
		// 		formatter.JSON(w, http.StatusOK, item)
		// 		return
		// 	}
		// }
	}
}

func createPayments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		// var result bson.M

		var payment Payment
		_ = json.NewDecoder(req.Body).Decode(&payment)

		uuid, _ := uuid.NewV4()
		payment.PaymentID = uuid.String()
		t := time.Now()
		payment.PaymentDate = t.Format("2006-01-02 15:04:05")
		payment.Status = true

		// err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
		err = c.Insert(payment)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Create new payment:", payment)
		formatter.JSON(w, http.StatusOK, payment)

		// fmt.Println("createPayments")
		// // params := mux.Vars(req)
		// var payment Payment
		// _ = json.NewDecoder(req.Body).Decode(&payment)

		// uuid, _ := uuid.NewV4()
		// payment.PaymentID = uuid.String()
		// t := time.Now()
		// payment.PaymentDate = t.Format("2006-01-02 15:04:05")
		// payment.Status = true

		// payments = append(payments, payment)

		// // fmt.Fprintf(w, "createPayments")
		// // json.NewEncoder(w).Encode(payments)
		// formatter.JSON(w, http.StatusOK, payments)
	}
}

func deletePayment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// session, err := mgo.Dial(mongodb_server)
		//     if err != nil {
		//             panic(err)
		//     }
		//     defer session.Close()
		//     session.SetMode(mgo.Monotonic, true)
		//     c := session.DB(mongodb_database).C(mongodb_collection)
		//     var result bson.M
		//     err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
		//     if err != nil {
		//             log.Fatal(err)
		//     }
		//     fmt.Println("Gumball Machine:", result )
		// formatter.JSON(w, http.StatusOK, result)

		fmt.Println("deletePayment")
		params := mux.Vars(req)
		for index, item := range payments {
			if item.PaymentID == params["id"] {
				payments = append(payments[:index], payments[index+1:]...)
				break
			}
		}

		// json.NewEncoder(w).Encode(payments)
		formatter.JSON(w, http.StatusOK, payments)
	}
}

func editPayment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// session, err := mgo.Dial(mongodb_server)
		//     if err != nil {
		//             panic(err)
		//     }
		//     defer session.Close()
		//     session.SetMode(mgo.Monotonic, true)
		//     c := session.DB(mongodb_database).C(mongodb_collection)
		//     var result bson.M
		//     err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
		//     if err != nil {
		//             log.Fatal(err)
		//     }
		//     fmt.Println("Gumball Machine:", result )
		// formatter.JSON(w, http.StatusOK, result)

		fmt.Println("editPayment")

		var payment Payment
		_ = json.NewDecoder(req.Body).Decode(&payment)

		fmt.Println("before", payment)

		params := mux.Vars(req)
		for index, item := range payments {
			if item.PaymentID == params["id"] {
				payment.PaymentID = item.PaymentID
				payment.UserID = item.UserID
				payment.OrderID = item.OrderID
				payment.PaymentDate = item.PaymentDate

				fmt.Println(payment)

				payments = append(payments[:index], payments[index+1:]...)
				payments = append(payments, payment)

				break
			}
		}

		// json.NewEncoder(w).Encode(payments)
		formatter.JSON(w, http.StatusOK, payments)
	}
}

// func getPaymentById(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("getPaymentByPaymentId")
// 	params := mux.Vars(req)
// 	for _, item := range payments {
// 		if item.PaymentID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}

// 	// fmt.Fprintf(w, "getPaymentByPaymentId")
// 	json.NewEncoder(w).Encode(&Payment{})
// }

// func createPayments(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("createPayments")
// 	// params := mux.Vars(req)
// 	var payment Payment
// 	_ = json.NewDecoder(req.Body).Decode(&payment)

// 	uuid, _ := uuid.NewV4()
// 	payment.PaymentID = uuid.String()
// 	t := time.Now()
// 	payment.PaymentDate = t.Format("2006-01-02 15:04:05")
// 	payment.Status = true

// 	payments = append(payments, payment)

// 	// fmt.Fprintf(w, "createPayments")
// 	json.NewEncoder(w).Encode(payments)
// }

// func deletePayment(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("deletePayment")
// 	params := mux.Vars(req)
// 	for index, item := range payments {
// 		if item.PaymentID == params["id"] {
// 			payments = append(payments[:index], payments[index+1:]...)
// 			break
// 		}
// 	}

// 	json.NewEncoder(w).Encode(payments)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Heyyooo")
// }

/*


 db.payments.insert({
	 PaymentID: '1',
	 UserID: '1',
	 OrderID: '100',
	 TotalAmount: NumberDecimal(100.50),
	 Status: true,
	 PaymentDate: Date('2018-11-11 20:27:43')
 });
*/
