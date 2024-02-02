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

	PORT := 8123
	ADDR := fmt.Sprintf("0.0.0.0:%d", PORT)

	fmt.Printf("The server is running at %v\n", ADDR)
	if err := http.ListenAndServe(ADDR, router); err != nil {
		log.Fatal(err)
	}

}
