// comment_controller.go

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MyGram/models"

	"github.com/gorilla/mux"
)

// CreateComment handles the creation of a new comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
    var newComment models.Comment
    err := json.NewDecoder(r.Body).Decode(&newComment)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = newComment.Save()
    if err != nil {
        http.Error(w, "Failed to create comment", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newComment)
}

// RetrieveCommentByID retrieves a comment by its ID
func RetrieveCommentByID(w http.ResponseWriter, r *http.Request) {
    commentIDStr := mux.Vars(r)["id"]
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    comment, err := models.GetCommentByID(uint(commentID))
    if err != nil {
        http.Error(w, "Comment not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(comment)
}

// UpdateCommentByID updates a comment by its ID
func UpdateCommentByID(w http.ResponseWriter, r *http.Request) {
    commentIDStr := mux.Vars(r)["id"]
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    var updatedComment models.Comment
    err = json.NewDecoder(r.Body).Decode(&updatedComment)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = models.UpdateCommentByID(uint(commentID), &updatedComment)
    if err != nil {
        http.Error(w, "Failed to update comment", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(updatedComment)
}

// DeleteCommentByID deletes a comment by its ID
func DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
    commentIDStr := mux.Vars(r)["id"]
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    err = models.DeleteCommentByID(uint(commentID))
    if err != nil {
        http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Comment deleted successfully"})
}
