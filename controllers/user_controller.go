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

// RegisterUser handles the registration of a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Hash the password before storing it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    newUser.Password = string(hashedPassword)

    // Save the user to the database
    err = newUser.Save()
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    // Respond with the newly registered user
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

// LoginUser handles user login
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

    // Find the user by email
    user, err := models.GetUserByEmail(loginData.Email)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Compare the hashed password with the provided password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "id":    user.ID,
    })

    // Sign the token with a secret key
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Respond with the JWT token
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// UpdateUser handles updating user information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from request parameters
    userIDStr := mux.Vars(r)["userId"]
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Decode request body into a user object
    var updatedUser models.User
    err = json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Get the user from the database
    user, err := models.GetUserByID(uint(userID))
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Update user information
    err = user.Update(&updatedUser)
    if err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    // Respond with updated user information
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser handles deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from request parameters
    userIDStr := mux.Vars(r)["userId"]
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Parse JWT token from request headers
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte((os.Getenv("JWT_SECRET_KEY"))), nil
    })
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    // Extract user ID from JWT claims
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

    // Ensure that the user ID from the token matches the requested user ID
    if userID != userIDFromToken {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Get the user from the database
    user, err := models.GetUserByID(uint(userID))
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Delete the user
    err = user.Delete()
    if err != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }

    // Respond with success message
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}