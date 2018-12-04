/*
	Gumball API in Go (Version 3)
	Uses MongoDB and RabbitMQ
	(For use with Kong API Key)
*/

package main

type Items struct {
	ItemId   string 	`json:"itemId"`
	ItemName string 	`json:"itemName"`
	Price 	 float32	`json:"price"`
	Description string  `json:"description"`
}	
type BurgerOrder struct {
	OrderId     string  `json:"orderId" bson:"orderId"`
	UserId      string  `json:"userId" bson:"userId"`
	OrderStatus string  `json:"OrderStatus" bson:"OrderStatus"`
	Cart        []Items `json:"Cart" bson:"Cart"`
	TotalAmount float32 `json:"TotalAmount" bson:"TotalAmount"`
}

type RequiredPayload struct {
	OrderId  string  	`json:"orderId" bson:"orderId"`
	ItemId   string 	`json:"itemId"`
	ItemName string 	`json:"itemName"`
	Price 	 float32	`json:"price"`
	Description string  `json:"description"`
}
var orders map[string]BurgerOrder
