/*
	Defining main function which is the entry point for the application.
	This is a Counter burger location module
	Refereneced from - https://github.com/paulnguyen/cmpe281/blob/master/golabs/godata/go-gumball-mongo/src/gumball/main.go
*/

package main

import (
	"os"
)

func main() {

	portNumber := os.Getenv("PORT")
	if len(portNumber) == 0 {
		portNumber = "3000"
	}

	server := NewServerConfiguration()
	server.Run(":" + portNumber)
}
