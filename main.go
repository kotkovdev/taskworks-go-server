package main

import (
	"fmt"
	"log"
	"net/http"
	"taskworks/app/routes"
)

func main() {
	fmt.Println("Running server...")
	fmt.Println("Server runned!")
	routes.Route()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
