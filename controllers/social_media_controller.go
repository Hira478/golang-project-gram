// social_media_controller.go

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MyGram/models"

	"github.com/gorilla/mux"
)

// CreateSocialMedia handles the creation of a new social media record
func CreateSocialMedia(w http.ResponseWriter, r *http.Request) {
    var newSocialMedia models.SocialMedia
    err := json.NewDecoder(r.Body).Decode(&newSocialMedia)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = newSocialMedia.Save()
    if err != nil {
        http.Error(w, "Failed to create social media", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newSocialMedia)
}

// RetrieveSocialMediaByID retrieves a social media record by its ID
func RetrieveSocialMediaByID(w http.ResponseWriter, r *http.Request) {
    // Extract social media ID from request parameters
    socialMediaIDStr := mux.Vars(r)["id"]
    socialMediaID, err := strconv.Atoi(socialMediaIDStr)
    if err != nil {
        http.Error(w, "Invalid social media ID", http.StatusBadRequest)
        return
    }

    socialMedia, err := models.GetSocialMediaByID(uint(socialMediaID))
    if err != nil {
        http.Error(w, "Social media not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(socialMedia)
}

// UpdateSocialMediaByID updates a social media record by its ID
func UpdateSocialMediaByID(w http.ResponseWriter, r *http.Request) {
    // Extract social media ID from request parameters
    socialMediaIDStr := mux.Vars(r)["id"]
    socialMediaID, err := strconv.Atoi(socialMediaIDStr)
    if err != nil {
        http.Error(w, "Invalid social media ID", http.StatusBadRequest)
        return
    }

    // Decode request body into a social media object
    var updatedSocialMedia models.SocialMedia
    err = json.NewDecoder(r.Body).Decode(&updatedSocialMedia)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Update social media information
    err = models.UpdateSocialMediaByID(uint(socialMediaID), &updatedSocialMedia)
    if err != nil {
        http.Error(w, "Failed to update social media", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(updatedSocialMedia)
}

// DeleteSocialMediaByID deletes a social media record by its ID
func DeleteSocialMediaByID(w http.ResponseWriter, r *http.Request) {
    // Extract social media ID from request parameters
    socialMediaIDStr := mux.Vars(r)["id"]
    socialMediaID, err := strconv.Atoi(socialMediaIDStr)
    if err != nil {
        http.Error(w, "Invalid social media ID", http.StatusBadRequest)
        return
    }

    // Delete the social media record
    err = models.DeleteSocialMediaByID(uint(socialMediaID))
    if err != nil {
        http.Error(w, "Failed to delete social media", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Social media deleted successfully"})
}
