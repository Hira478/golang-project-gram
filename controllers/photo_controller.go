// controllers/photo_controller.go

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MyGram/models"

	"github.com/gorilla/mux"
)

// CreatePhoto handles creating a new photo
func CreatePhoto(w http.ResponseWriter, r *http.Request) {
    var newPhoto models.Photo
    err := json.NewDecoder(r.Body).Decode(&newPhoto)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Save the photo to the database
    err = newPhoto.Save()
    if err != nil {
        http.Error(w, "Failed to create photo", http.StatusInternalServerError)
        return
    }

    // Respond with the newly created photo
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newPhoto)
}

// GetPhoto handles retrieving a photo by its ID
func GetPhoto(w http.ResponseWriter, r *http.Request) {
    // Extract photo ID from request parameters
    photoIDStr := mux.Vars(r)["photoId"]
    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Get the photo from the database
    photo, err := models.GetPhotoByID(uint(photoID))
    if err != nil {
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // Respond with the photo
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(photo)
}

// UpdatePhoto handles updating a photo
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
    // Implement the logic for updating a photo here
}

// DeletePhoto handles deleting a photo
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
    // Implement the logic for deleting a photo here
}
