// main.go

package main

import (
	"MyGram/models"
	"MyGram/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    // Initialize the database connection
    models.InitDB()

    // Initialize the router
    router := mux.NewRouter()

    // Register routes for users, photos, social media, and comments
    routes.SetupUserRoutes(router)
    routes.SetupPhotoRoutes(router)
    routes.SetupSocialMediaRoutes(router)
    routes.SetupCommentRoutes(router)

    // Start the HTTP server
    port := ":8080" // specify the port you want to listen on
    fmt.Printf("Server listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, router))
}
