package main

// imports
import (
	"fmt"
	"github.com/natnael-eyuel-dev/Task-Manager-REST-API/router"
)

// entry point of the Task Manager REST API application
func main() {
	fmt.Println("Task Manager REST API Project")    // print startup message

	router := router.SetupRouter()	   // initialize the router with all configured routes
	
	router.Run(":8080")      // start the server on port 8080
}
