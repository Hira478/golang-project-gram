// routes/photo_routes.go

package routes

import (
	"MyGram/controllers"

	"github.com/gorilla/mux"
)

// SetupPhotoRoutes sets up routes for photo-related endpoints
func SetupPhotoRoutes(router *mux.Router) {
    router.HandleFunc("/photos", controllers.CreatePhoto).Methods("POST")
    router.HandleFunc("/photos/{photoId}", controllers.GetPhoto).Methods("GET")
    router.HandleFunc("/photos/{photoId}", controllers.UpdatePhoto).Methods("PUT")
    router.HandleFunc("/photos/{photoId}", controllers.DeletePhoto).Methods("DELETE")
}
