package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
)

var people []User
var mongodb_server = "localhost"
var mongodb_database = "burger"
var mongodb_collection = "Users"

func GetUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var m User
    _ = json.NewDecoder(req.Body).Decode(&m)		
    fmt.Println("Get data of user: ", params["id"])
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
           panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"id" : params["id"]}
    var result bson.M
    err = c.Find(query).One(&result)
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("User:", result )
	json.NewEncoder(w).Encode(result)
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
           panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{}
    var result []bson.M
    err = c.Find(query).All(&result)
    if err != nil {
            log.Fatal(err)
    }
	fmt.Println("User:", result )
	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)
	unqueId := uuid.Must(uuid.NewV4())
	person.Id = unqueId.String()
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
           panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    err = c.Insert(person)
    if err != nil {
            log.Fatal(err)
    }
	json.NewEncoder(w).Encode(person)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
           panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"id":params["Id"]}
    err = c.Remove(query)
    if err != nil {
            log.Fatal(err)
    }
	json.NewEncoder(w).Encode(people)
}
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)
	params := mux.Vars(req)
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
           panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"id":params["id"]}
	updator := bson.M{
				"$set": bson.M{
						"firstname": person.Firstname,
						"lastname":person.Lastname,
						"address":person.Address,
						"password":person.Password}}
    err = c.Update(query, updator)
    if err != nil {
            log.Fatal(err)
    }
	json.NewEncoder(w).Encode(updator)
}
func UserSignIn(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range people {
		if item.Id == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/signup", CreateUser).Methods("POST")
	router.HandleFunc("/users/signin", UserSignIn).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}