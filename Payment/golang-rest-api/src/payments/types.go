/*
	Payment API in Go
*/

package main

type Payment struct {
	PaymentID   string  `json:"paymentId, mitempty"`
	UserID      string  `json:"userId,omitempty"`
	OrderID     string  `json:"orderId,omitempty"`
	TotalAmount float32 `json:"totalAmount,omitempty"`
	Status      bool    `json:"status,omitempty"`
	PaymentDate string  `json:"paymentDate,omitempty"`
}
