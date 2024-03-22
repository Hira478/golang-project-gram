// comment_routes.go

package routes

import (
	"MyGram/controllers"

	"github.com/gorilla/mux"
)

// SetupCommentRoutes sets up routes for comment CRUD operations
func SetupCommentRoutes(router *mux.Router) {
    router.HandleFunc("/comments", controllers.CreateComment).Methods("POST")
    router.HandleFunc("/comments/{id}", controllers.RetrieveCommentByID).Methods("GET")
    router.HandleFunc("/comments/{id}", controllers.UpdateCommentByID).Methods("PUT")
    router.HandleFunc("/comments/{id}", controllers.DeleteCommentByID).Methods("DELETE")
}
