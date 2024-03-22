// social_media_controller.go

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MyGram/models"

	"github.com/gorilla/mux"
)

// CreateSocialMedia menangani pembuatan catatan media sosial baru
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

// RetrieveSocialMediaByID mengambil catatan media sosial dengan ID-nya
func RetrieveSocialMediaByID(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID media sosial dari parameter permintaan
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

// UpdateSocialMediaByID memperbarui catatan media sosial dengan ID-nya
func UpdateSocialMediaByID(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID media sosial dari parameter permintaan
    socialMediaIDStr := mux.Vars(r)["id"]
    socialMediaID, err := strconv.Atoi(socialMediaIDStr)
    if err != nil {
        http.Error(w, "Invalid social media ID", http.StatusBadRequest)
        return
    }

    // Memecahkan kode isi permintaan menjadi objek media sosial
    var updatedSocialMedia models.SocialMedia
    err = json.NewDecoder(r.Body).Decode(&updatedSocialMedia)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Perbarui informasi media sosial
    err = models.UpdateSocialMediaByID(uint(socialMediaID), &updatedSocialMedia)
    if err != nil {
        http.Error(w, "Failed to update social media", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(updatedSocialMedia)
}

// DeleteSocialMediaByID menghapus catatan media sosial dengan ID-nya
func DeleteSocialMediaByID(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID media sosial dari parameter permintaan
    socialMediaIDStr := mux.Vars(r)["id"]
    socialMediaID, err := strconv.Atoi(socialMediaIDStr)
    if err != nil {
        http.Error(w, "Invalid social media ID", http.StatusBadRequest)
        return
    }

    // Menghapus rekaman media sosial
    err = models.DeleteSocialMediaByID(uint(socialMediaID))
    if err != nil {
        http.Error(w, "Failed to delete social media", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Social media deleted successfully"})
}
