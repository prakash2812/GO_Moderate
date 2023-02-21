package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/arjun/modules/25-mongoapi/router"
)

func main() {
	fmt.Println("ModoDB setup")
	route := router.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":3000", route))
}
