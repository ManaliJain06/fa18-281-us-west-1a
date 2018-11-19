package db

import mgo "gopkg.in/mgo.v2"

var mongoHostName = "localhost"
var mongoDatabase = "BurgerUsers"
var mongoCollection = "Users"
var MgoSession *mgo.Session

 // export the mongo session

 func init() {
	 session, err := mgo.Dial(mongoHostName)
	 if err != nil {
		 panic(err)
	 }
	 MgoSession = session
 }