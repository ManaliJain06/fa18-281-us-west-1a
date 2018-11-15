/*
	Gumball API in Go (Version 3)
	Uses MongoDB and RabbitMQ
	(For use with Kong API Key)
*/

package main

type items struct {
	itemId   string
	quantity int
}
type burgerOrder struct {
	OrderId     string
	UserId      string
	OrderStatus string
	TotalAmount float32
}

var orders map[string]burgerOrder
