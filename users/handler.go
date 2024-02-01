package users

import (
	"fmt"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get user details
	fmt.Fprintf(w, "Getting user details")
}
