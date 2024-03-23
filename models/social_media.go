// social_media.go

package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// ErrSocialMediaNotFound kesalahan yang dikembalikan ketika catatan media sosial tidak ditemukan
var ErrSocialMediaNotFound = errors.New("social media not found")

// SocialMedia mewakili entitas dari media sosial
type SocialMedia struct {
    ID             uint   `gorm:"primary_key"`
    Name           string `gorm:"not null"`
    SocialMediaURL string `gorm:"not null"`
    UserID         uint   `gorm:"not null"`
    CreatedAt      time.Time
    UpdatedAt      time.Time
}

func (s *SocialMedia) TableName() string {
    return "my_schema.social_media"
}

// Simpan catatan media sosial baru ke dalam database
func (sm *SocialMedia) Save() error {
    if err := DB.Create(sm).Error; err != nil {
        return err
    }
    return nil
}

// GetSocialMediaByID mengambil catatan media sosial dengan ID-nya
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

// UpdateSocialMediaByID memperbarui catatan media sosial dengan ID-nya
func UpdateSocialMediaByID(id uint, updatedSM *SocialMedia) error {
    if err := DB.Model(&SocialMedia{}).Where("id = ?", id).Updates(updatedSM).Error; err != nil {
        return err
    }
    return nil
}

// DeleteSocialMediaByID menghapus catatan media sosial dengan ID-nya
func DeleteSocialMediaByID(id uint) error {
    if err := DB.Where("id = ?", id).Delete(&SocialMedia{}).Error; err != nil {
        return err
    }
    return nil
}
