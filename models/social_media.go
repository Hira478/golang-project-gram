// social_media.go

package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// ErrSocialMediaNotFound is the error returned when a social media record is not found
var ErrSocialMediaNotFound = errors.New("social media not found")

// SocialMedia represents a social media entity
type SocialMedia struct {
    ID             uint   `gorm:"primary_key"`
    Name           string `gorm:"not null"`
    SocialMediaURL string `gorm:"not null"`
    UserID         uint   `gorm:"not null"`
    CreatedAt      time.Time
    UpdatedAt      time.Time
}

// Save inserts a new social media record into the database
func (sm *SocialMedia) Save() error {
    if err := DB.Create(sm).Error; err != nil {
        return err
    }
    return nil
}

// GetSocialMediaByID retrieves a social media record by its ID
func GetSocialMediaByID(id uint) (*SocialMedia, error) {
    var sm SocialMedia
    if err := DB.First(&sm, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            return nil, ErrSocialMediaNotFound
        }
        return nil, err
    }
    return &sm, nil
}

// UpdateSocialMediaByID updates a social media record by its ID
func UpdateSocialMediaByID(id uint, updatedSM *SocialMedia) error {
    if err := DB.Model(&SocialMedia{}).Where("id = ?", id).Updates(updatedSM).Error; err != nil {
        return err
    }
    return nil
}

// DeleteSocialMediaByID deletes a social media record by its ID
func DeleteSocialMediaByID(id uint) error {
    if err := DB.Where("id = ?", id).Delete(&SocialMedia{}).Error; err != nil {
        return err
    }
    return nil
}
