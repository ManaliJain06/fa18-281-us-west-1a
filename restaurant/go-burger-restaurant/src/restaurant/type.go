package main

type restaurant struct { 
	RestaurantName 	string	`json:"restaurantName"`
	RestaurantId 	string	`json:"restaurantId"`
	Zipcode			string  `json:"zipcode"`
	Phone			string	`json:"phone"`
	AddressLine1	string  `json:"addressLine1"`
	AddressLine2	string  `json:"addressLine2"`
	City			string  `json:"city"`
	State			string  `json:"state"`
	Country			string  `json:"country"`
	Hours			string  `json:"hours"`
	AcceptedCards	string  `json:"acceptedCards"`
	Distance		string  `json:"distance"`
	Email			string  `json:"email"`
}

    
