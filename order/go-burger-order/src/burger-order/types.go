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
	OrderStatus string  `json:"orderStatus" bson:"orderStatus"`
	Cart        []Items `json:"items" bson:"items"`
	TotalAmount float32 `json:"totalAmount" bson:"totalAmount"`
	IpAddress	string	`json:"ipaddress" bson:"ipaddress"`
}

type RequiredPayload struct {
	OrderId  string  	`json:"orderId" bson:"orderId"`
	UserId   string  	`json:"userId" bson:"userId"`
	ItemId   string 	`json:"itemId"`
	ItemName string 	`json:"itemName"`
	Price 	 float32	`json:"price"`
	Description string  `json:"description"`
}
var orders map[string]BurgerOrder
