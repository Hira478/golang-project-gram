package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"MyGram/models"
)

// RegisterUser menangani pendaftaran pengguna baru
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Hash kata sandi sebelum menyimpannya
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    newUser.Password = string(hashedPassword)

    // Menyimpan pengguna ke database
    err = newUser.Save()
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    // Tanggapi dengan pengguna yang baru terdaftar
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

// LoginUser menangani login pengguna
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    err := json.NewDecoder(r.Body).Decode(&loginData)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Menemukan pengguna melalui email
    user, err := models.GetUserByEmail(loginData.Email)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Bandingkan kata sandi hash dengan kata sandi yang diberikan
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Hasilkan token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "id":    user.ID,
    })

    // Tanda tangani token dengan kunci rahasia
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Merespons dengan token JWT
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// UpdateUser menangani pembaruan informasi pengguna
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID pengguna dari parameter permintaan
    userIDStr := mux.Vars(r)["userId"]
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Mendekode isi permintaan menjadi objek pengguna
    var updatedUser models.User
    err = json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Dapatkan pengguna dari database
    user, err := models.GetUserByID(uint(userID))
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Perbarui informasi pengguna
    err = user.Update(&updatedUser)
    if err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    // Merespons dengan informasi pengguna yang diperbarui
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser menangani penghapusan pengguna
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Ekstrak ID pengguna dari parameter permintaan
    userIDStr := mux.Vars(r)["userId"]
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Mengurai token JWT dari header permintaan
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte((os.Getenv("JWT_SECRET_KEY"))), nil
    })
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    // Mengekstrak ID pengguna dari klaim JWT
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    userIDFromToken, err := strconv.Atoi(claims["id"].(string))
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    // Pastikan ID pengguna dari token cocok dengan ID pengguna yang diminta
    if userID != userIDFromToken {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Dapatkan pengguna dari database
    user, err := models.GetUserByID(uint(userID))
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Menghapus pengguna
    err = user.Delete()
    if err != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }

    // Tanggapi dengan pesan sukses
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}