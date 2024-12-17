//debugging by chatGPT
//LINK: https://chatgpt.com/share/676172a5-53a0-8002-a410-df12ef486f29
// structure is as follows 
/*
04mongoDB
│
├── controllers/
│   └── controller.go  # Contains all controller functions for CRUD operations
│
├── models/
│   └── model.go       # Contains the Netflix struct model for the movie data
│
├── router/
│   └── router.go      # Contains all route handlers
│
├── main.go            # Main entry point of the application to start the server
├── go.mod             # Go modules file
├── go.sum   
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/The-Fermits/Golang/router"
)

func main() {
	fmt.Println("Hey we will here see how to use mongoDB in golang")
	fmt.Println("wait server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("serving/listening at port 4000")
}
