/*

 */
 package main

 import (
	 "encoding/json"
	 "fmt"
	 "log"
	 "net"
	 "net/http"
	 "os"
	 "time"
 
	 "strings"
 
	 "github.com/codegangsta/negroni"
	 "github.com/gorilla/handlers"
	 "github.com/gorilla/mux"
	 uuid "github.com/satori/go.uuid"
	 "github.com/unrolled/render"
	 mgo "gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
 )
 
 /*
	 Mac commands to start and stop local mongo
	 ===============================================
	 brew services start mongodb
	 brew services stop mongodb
	 brew services restart mongodb
 */
 
 // MongoDB Config
 // var mongodb_server = "localhost:27017"
 // var mongodb_database = "test" // cmpe281
 // var mongodb_collection = "payments"
 
 // Use EC2 MongoDB Sharding, get values from environment variables
 var mongodb_server = os.Getenv("AWS_MONGODB")
 var mongodb_database = os.Getenv("MONGODB_DBNAME")
 var mongodb_collection = os.Getenv("MONGODB_COLLECTION")
 var mongodb_username = os.Getenv("MONGODB_USERNAME")
 var mongodb_password = os.Getenv("MONGODB_PASSWORD")
 
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

	 allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	 allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	 allowedOrigins := handlers.AllowedOrigins([]string{"*"})
 
	 n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(mx))
 
	 return n
 }
 
 // API Routes
 func initRoutes(mx *mux.Router, formatter *render.Render) {
	 t := time.Now()
	 payments = append(payments, Payment{PaymentID: "1", UserID: "1", OrderID: "11", TotalAmount: 100.50, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})
	 payments = append(payments, Payment{PaymentID: "2", UserID: "2", OrderID: "22", TotalAmount: 30.30, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})
 
	 mx.HandleFunc("/payments/ping", pingHandler(formatter)).Methods("GET")
	 mx.HandleFunc("/payments", getAllPayments(formatter)).Methods("GET")
	 mx.HandleFunc("/payments/{id}", getPaymentByID(formatter)).Methods("GET")
	 mx.HandleFunc("/payments", createPayments(formatter)).Methods("POST")
	 mx.HandleFunc("/payments/{id}", deletePayment(formatter)).Methods("DELETE")
	 mx.HandleFunc("/payments/{id}", editPayment(formatter)).Methods("PUT")
 }
 
 func handleRequest() {
 
	 t := time.Now()
	 payments = append(payments, Payment{PaymentID: "1", UserID: "1", OrderID: "11", TotalAmount: 100.50, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})
	 payments = append(payments, Payment{PaymentID: "2", UserID: "2", OrderID: "22", TotalAmount: 30.30, Status: true, PaymentDate: t.Format("2006-01-02 15:04:05")})
 }
 
 // API Ping Handler
 func pingHandler(formatter *render.Render) http.HandlerFunc {
	 return func(w http.ResponseWriter, req *http.Request) {
		 message := "Burger Payments API Server Working on machine: " + getSystemIp()
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
  
 func getAllPayments(formatter *render.Render) http.HandlerFunc {
	 return func(w http.ResponseWriter, req *http.Request) {
		 session, _ := mgo.Dial(mongodb_server)
		 err := session.DB("admin").Login(mongodb_username, mongodb_password)
		 if err != nil {
			 formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			 return
		 }
		 defer session.Close()
		 session.SetMode(mgo.Monotonic, true)
		 c := session.DB(mongodb_database).C(mongodb_collection)
		 var result []bson.M
		 err = c.Find(nil).All(&result)
		 if err != nil {
			 formatter.JSON(w, http.StatusNotFound, "Get All Payment Error")
			 return
		 }
		 fmt.Println("getAllPayments:", result)
		 formatter.JSON(w, http.StatusOK, result)
	 }
 }
 
 func getPaymentByID(formatter *render.Render) http.HandlerFunc {
	 return func(w http.ResponseWriter, req *http.Request) {
		 session, _ := mgo.Dial(mongodb_server)
		 err := session.DB("admin").Login(mongodb_username, mongodb_password)
		 if err != nil {
			 formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			 return
		 }
		 defer session.Close()
		 session.SetMode(mgo.Monotonic, true)
		 c := session.DB(mongodb_database).C(mongodb_collection)
		 var result bson.M
		 params := mux.Vars(req)
 
		 fmt.Printf("params[id]=%s \n", params["id"])
 
		 err = c.Find(bson.M{"paymentid": params["id"]}).One(&result)
		 if err != nil {
			 formatter.JSON(w, http.StatusNotFound, "Get Payment by ID Error")
			 return
		 }
		 fmt.Println("getPaymentByID:", result)
		 formatter.JSON(w, http.StatusOK, result)
	 }
 }
 
 func createPayments(formatter *render.Render) http.HandlerFunc {
	 return func(w http.ResponseWriter, req *http.Request) {
		 session, _ := mgo.Dial(mongodb_server)
		 err := session.DB("admin").Login(mongodb_username, mongodb_password)
		 if err != nil {
			 formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			 return
		 }
		 defer session.Close()
		 session.SetMode(mgo.Monotonic, true)
		 c := session.DB(mongodb_database).C(mongodb_collection)
 
		 var payment Payment
		 _ = json.NewDecoder(req.Body).Decode(&payment)
 
		 uuid, _ := uuid.NewV4()
		 payment.PaymentID = uuid.String()
		 t := time.Now()
		 payment.PaymentDate = t.Format("2006-01-02 15:04:05")
		 payment.Status = true
 
		 err = c.Insert(payment)
		 if err != nil {
			 formatter.JSON(w, http.StatusNotFound, "Create Payment Error")
			 return
		 }
		 fmt.Println("Create new payment:", payment)
		 formatter.JSON(w, http.StatusOK, payment)
	 }
 }
 
 func deletePayment(formatter *render.Render) http.HandlerFunc {
	 return func(w http.ResponseWriter, req *http.Request) {
		 session, _ := mgo.Dial(mongodb_server)
		 err := session.DB("admin").Login(mongodb_username, mongodb_password)
		 if err != nil {
			 formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			 return
		 }
		 defer session.Close()
		 session.SetMode(mgo.Monotonic, true)
		 c := session.DB(mongodb_database).C(mongodb_collection)
		 var result Payment
		 params := mux.Vars(req)
 
		 err = c.Find(bson.M{"paymentid": params["id"]}).One(&result)
		 if err != nil {
			 formatter.JSON(w, http.StatusNotFound, "Delete Payment: cannot find ID Error")
			 return
		 } else {
			 fmt.Println("result.PaymentID :" + result.PaymentID)
			 fmt.Println("result.OrderID   :" + result.OrderID)
 
			 err = c.Remove(bson.M{"paymentid": result.PaymentID})
			 if err != nil {
				 fmt.Println("error:" + err.Error())
				 formatter.JSON(w, http.StatusNotFound, "Delete Payment: delete Error")
				 return
			 }
		 }
		 fmt.Println("deletePayment ", result)
		 formatter.JSON(w, http.StatusOK, result)
	 }
 }
 
 func editPayment(formatter *render.Render) http.HandlerFunc {
	 return func(w http.ResponseWriter, req *http.Request) {
 
		 // find payment, then edit
		 var payment Payment
		 _ = json.NewDecoder(req.Body).Decode(&payment)
		 fmt.Println("Edit payment to these attributes: ", payment)
		 session, _ := mgo.Dial(mongodb_server)
		 err := session.DB("admin").Login(mongodb_username, mongodb_password)
		 if err != nil {
			 formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			 return
		 }
		 defer session.Close()
		 session.SetMode(mgo.Monotonic, true)
		 c := session.DB(mongodb_database).C(mongodb_collection)
		 params := mux.Vars(req)
 
		 var fetchedPayment bson.M
		 err = c.Find(bson.M{"paymentid": params["id"]}).One(&fetchedPayment)
		 if err != nil {
			 formatter.JSON(w, http.StatusNotFound, "Edit Payment Error")
			 return
		 }
 
		 fmt.Println("fetchedPayment", fetchedPayment)
		 fetchedPayment["totalamount"] = payment.TotalAmount
		 fetchedPayment["status"] = payment.Status
 
		 query := bson.M{"paymentid": params["id"]}
		 err = c.Update(query, &fetchedPayment)
		 if err != nil {
			 log.Fatal(err)
		 }
 
		 fmt.Println("Edit Payment:", fetchedPayment)
		 formatter.JSON(w, http.StatusOK, fetchedPayment)
	 }
 }
 

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
 