/*
	Burger users API
	Port mapping and initiate the user services
*/

package main

import (
	"os"
	"github.com/kabukky/httpscerts"
	"log"
)


func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	// Check if the cert files are available.
    err := httpscerts.Check("cert.pem", "key.pem")
    // If they are not available, generate new ones.
    if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8000")
        if err != nil {
            log.Fatal("Error: Couldn't create https certs.")
        }
    }

	server := MenuServer()
	server.RunTLS(":" + port,"cert.pem", "key.pem")
}