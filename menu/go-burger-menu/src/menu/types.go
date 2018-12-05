/*
	Burger Menu Item API
*/
	
package main

// Item structure 
type Item struct {

	Id             	string 	`json:"id"`
	Name		    string   `json:"name"` 	
	Price 			int	    `json:"price"`
	Description 	string	`json:"description"`
	Calories		int     `json:"calories"`
}

type Menu struct {
	RestaurantId	string	`json:"resId"`
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
	Item			Item   `json:"item"`
}

type deleteReqBody struct {
	RestaurantId	string	`json:"resId"`
	ItemId			string `json:"itemId`
}


