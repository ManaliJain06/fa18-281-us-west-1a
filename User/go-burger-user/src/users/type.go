package main

type User struct {
	Id        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
}

type Address struct {
	Street string `json:"street,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Zipcode string `json:"zipcode,omitempty:`
}