/*
	Burger Menu Item API
*/
	
package main

// Item structure 
type Item struct {

	Id             	string 	
	Name		    string    	
	Price 			int	    
	Description 	string	
	Calories		int
}

type Menu struct {
	RestaurantId	string	`json:"resId"`
	RestaurantName	string 	
	Items 			[]Item   `json:"items"`
}

type menuItem struct {

	Id             	string 	
	Name		    string    	
	Price 			int	    
	Description 	string	
	Calories		int
}

type restaurantReqBody struct {
	RestaurantId	string	`json:"resId"`
	RestaurantName	string 	
	Item			Item   `json:"item"`
}


