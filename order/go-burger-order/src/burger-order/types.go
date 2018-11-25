/*
	Gumball API in Go (Version 3)
	Uses MongoDB and RabbitMQ
	(For use with Kong API Key)
*/

package main

type items struct {
	ItemId   string
	Quantity int
}
type burgerOrder struct {
	OrderId     string  `bson:"OrderId"`
	UserId      string  `bson:"UserId"`
	OrderStatus string  `bson:"OrderStatus"`
	TotalAmount float32 `bson:"TotalAmount"`
	Cart        []items `bson:"Cart"`
}

var orders map[string]burgerOrder
