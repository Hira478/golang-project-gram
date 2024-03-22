// social_media_routes.go

package routes

import (
	"MyGram/controllers"

	"github.com/gorilla/mux"
)

// SetupSocialMediaRoutes sets up routes for social media CRUD operations
func SetupSocialMediaRoutes(router *mux.Router) {
    router.HandleFunc("/social-media", controllers.CreateSocialMedia).Methods("POST")
    router.HandleFunc("/social-media/{id}", controllers.RetrieveSocialMediaByID).Methods("GET")
    router.HandleFunc("/social-media/{id}", controllers.UpdateSocialMediaByID).Methods("PUT")
    router.HandleFunc("/social-media/{id}", controllers.DeleteSocialMediaByID).Methods("DELETE")
}
