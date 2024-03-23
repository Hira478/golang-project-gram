package main

import (
	"MyGram/models"
	"MyGram/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Memuuat variabel lingkungan dari file .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Menginisialisasi koneksi database
    models.InitDB()

    // Menginisialisasi router
    router := mux.NewRouter()

    // register rute untuk pengguna, foto, media sosial, dan komentar
    routes.SetupUserRoutes(router)
    routes.SetupPhotoRoutes(router)
    routes.SetupSocialMediaRoutes(router)
    routes.SetupCommentRoutes(router)

    // Mulai server HTTP
    port := ":8080"
    fmt.Printf("Server listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, router))
}
