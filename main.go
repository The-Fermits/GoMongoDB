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
