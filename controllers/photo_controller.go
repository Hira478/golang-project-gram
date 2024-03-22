// controllers/photo_controller.go

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MyGram/models"

	"github.com/gorilla/mux"
)

// CreatePhoto menangani pembuatan foto baru
func CreatePhoto(w http.ResponseWriter, r *http.Request) {
    var newPhoto models.Photo
    err := json.NewDecoder(r.Body).Decode(&newPhoto)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Simpan foto ke database
    err = newPhoto.Save()
    if err != nil {
        http.Error(w, "Failed to create photo", http.StatusInternalServerError)
        return
    }

    // Tanggapi dengan foto yang baru dibuat
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newPhoto)
}

// GetPhoto menangani pengambilan foto dengan ID-nya
func GetPhoto(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID foto dari parameter permintaan
    photoIDStr := mux.Vars(r)["photoId"]
    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Dapatkan foto dari database
    photo, err := models.GetPhotoByID(uint(photoID))
    if err != nil {
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // Tanggapi dengan foto
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(photo)
}

// UpdatePhoto menangani pembaruan foto
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID foto dari parameter permintaan
    photoIDStr := mux.Vars(r)["photoId"]
    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Mengubah request body menjadi objek foto yang diperbarui
    var updatedPhoto models.Photo
    err = json.NewDecoder(r.Body).Decode(&updatedPhoto)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Dapatkan foto dari database
    photo, err := models.GetPhotoByID(uint(photoID))
    if err != nil {
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // Perbarui informasi foto
    err = photo.Update(&updatedPhoto)
    if err != nil {
        http.Error(w, "Failed to update photo", http.StatusInternalServerError)
        return
    }

    // Tanggapi dengan foto yang diperbarui
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedPhoto)
}

// DeletePhoto menangani penghapusan foto
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID foto dari parameter permintaan
    photoIDStr := mux.Vars(r)["photoId"]
    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Dapatkan foto dari database
    photo, err := models.GetPhotoByID(uint(photoID))
    if err != nil {
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // Menghapus foto
    err = photo.Delete()
    if err != nil {
        http.Error(w, "Failed to delete photo", http.StatusInternalServerError)
        return
    }

    // Tanggapi dengan pesan sukses
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Photo deleted successfully"})
}
