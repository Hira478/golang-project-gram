// photo.go

package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Photo struct {
    ID        uint   `gorm:"primary_key"`
    Title     string `gorm:"not null"`
    Caption   string
    PhotoURL  string `gorm:"not null"`
    UserID    uint
    CreatedAt time.Time
    UpdatedAt time.Time
}

// ErrPhotoNotFound is the error returned when a photo is not found
var ErrPhotoNotFound = errors.New("photo not found")

// Save saves a new photo record to the database
func (p *Photo) Save() error {
    if err := DB.Create(&p).Error; err != nil {
        return err
    }
    return nil
}

// GetPhotoByID retrieves a photo by its ID from the database
func GetPhotoByID(photoID uint) (*Photo, error) {
    var photo Photo
    if err := DB.First(&photo, photoID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, ErrPhotoNotFound
        }
        return nil, err
    }
    return &photo, nil
}

// Update updates an existing photo record in the database
func (p *Photo) Update(updatedPhoto *Photo) error {
    if err := DB.Model(p).Updates(updatedPhoto).Error; err != nil {
        return err
    }
    return nil
}

// Delete deletes a photo record from the database
func (p *Photo) Delete() error {
    if err := DB.Delete(&p).Error; err != nil {
        return err
    }
    return nil
}