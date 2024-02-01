package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ismail-alokin/server-sample-1/users"
)

func main() {
	log.SetFlags(log.Lshortfile)

	PORT := 8081
	ADDR := fmt.Sprintf("localhost:%d", PORT)

	router := http.NewServeMux()

	router.HandleFunc("/api/hello", helloHandler)
	router.HandleFunc("/api/user", users.GetUserHandler)

	fmt.Printf("Server running at %v\n", ADDR)
	if err := http.ListenAndServe(ADDR, router); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Someone hit the hello handler")
	fmt.Fprintf(w, "Hello, World!")
}
