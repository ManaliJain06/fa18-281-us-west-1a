/*
	Burger Menu Items API
	Port mapping and entry of menu API services
*/

package main

import (
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := MenuServer()
	server.Run(":" + port)
}
