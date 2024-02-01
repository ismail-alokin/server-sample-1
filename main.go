package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ismail-alokin/server-sample-1/pkg/routes"
)

func main() {

	log.SetFlags(log.Lshortfile)
	router := routes.SetupRoutes()

	PORT := 8081
	ADDR := fmt.Sprintf("localhost:%d", PORT)

	fmt.Printf("Server running at %v\n", ADDR)
	if err := http.ListenAndServe(ADDR, router); err != nil {
		log.Fatal(err)
	}

}
