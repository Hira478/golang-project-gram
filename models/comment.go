// comment.go

package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// ErrCommentNotFound is the error returned when a comment is not found
var ErrCommentNotFound = errors.New("comment not found")

// Comment represents a comment entity
type Comment struct {
    ID        uint   `gorm:"primary_key"`
    UserID    uint   `gorm:"not null"`
    PhotoID   uint   `gorm:"not null"`
    Message   string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Save inserts a new comment record into the database
func (c *Comment) Save() error {
    if err := DB.Create(c).Error; err != nil {
        return err
    }
    return nil
}

// GetCommentByID retrieves a comment record by its ID
func GetCommentByID(id uint) (*Comment, error) {
    var comment Comment
    if err := DB.First(&comment, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            return nil, ErrCommentNotFound
        }
        return nil, err
    }
    return &comment, nil
}

// UpdateCommentByID updates a comment record by its ID
func UpdateCommentByID(id uint, updatedComment *Comment) error {
    if err := DB.Model(&Comment{}).Where("id = ?", id).Updates(updatedComment).Error; err != nil {
        return err
    }
    return nil
}

// DeleteCommentByID deletes a comment record by its ID
func DeleteCommentByID(id uint) error {
    if err := DB.Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
        return err
    }
    return nil
}
