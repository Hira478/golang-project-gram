// routes/user_routes.go

package routes

import (
	"MyGram/controllers"

	"github.com/gorilla/mux"
)

// Set up routes for user-related endpoints
func SetupUserRoutes(router *mux.Router) {
    router.HandleFunc("/users/register", controllers.RegisterUser).Methods("POST")
    router.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")
    router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE") // New route for deleting user
}
