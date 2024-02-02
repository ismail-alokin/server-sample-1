package routes

import (
	"github.com/gorilla/mux"
	"github.com/ismail-alokin/server-sample-1/pkg/users"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", users.GetAllUsersHandler)
	router.HandleFunc("/api/users/html", users.GetUserHtml)
	router.HandleFunc("/api/users/{user_id}", users.GetUserHandler)

	return router
}
