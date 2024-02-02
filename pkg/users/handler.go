package users

import (
	"fmt"
	"log"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUserHandler")
	fmt.Fprintf(w, "Getting user details")
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllUsersHandler")
	fmt.Fprintf(w, "Getting all users")
}

func GetUserHtml(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUserHtml")
	fmt.Fprintf(w, "Getting ")
}
