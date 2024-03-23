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

func (p *Photo) TableName() string {
    return "my_schema.photos"
}

// ErrPhotoNotFound adalah kesalahan yang dikembalikan ketika foto tidak ditemukan
var ErrPhotoNotFound = errors.New("photo not found")

// Simpan menyimpan catatan foto baru ke database
func (p *Photo) Save() error {
    if err := DB.Create(&p).Error; err != nil {
        return err
    }
    return nil
}

// GetPhotoByID mengambil foto dengan ID-nya dari database
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

// Pembaruan foto yang ada di database
func (p *Photo) Update(updatedPhoto *Photo) error {
    if err := DB.Model(p).Updates(updatedPhoto).Error; err != nil {
        return err
    }
    return nil
}

// Hapus foto dari database
func (p *Photo) Delete() error {
    if err := DB.Delete(&p).Error; err != nil {
        return err
    }
    return nil
}