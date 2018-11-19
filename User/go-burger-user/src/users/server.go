package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
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
    query := bson.M{"Id" : params["id"]}
    var result bson.M
    err = c.Find(query).One(&result)
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("User:", result )
	json.NewEncoder(w).Encode(result)
	// for _, item := range people {
	// 	if item.Id == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Person{})
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person User
	params := mux.Vars(req)
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
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
func UpdateUser(w http.ResponseWriter, req *http.Request) {
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
	// people = append(people, Person{Id: "1", Firstname: "varun", Lastname: "jindal", Address: &Address{Street:"350 e taylo street",City: "Delhi", State: "Delhi",Zipcode:"95112"}, Email:"varun.jindal@sjus.edu", Password:"varun"})
	// people = append(people, Person{Id: "2", Firstname: "Manali", Lastname: "Jain"})
	router.HandleFunc("/users", GetAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/signup", CreateUser).Methods("POST")
	router.HandleFunc("/users/signin", UserSignIn).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}