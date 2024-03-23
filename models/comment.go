// comment.go

package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// ErrCommentNotFound adalah kesalahan yang dikembalikan ketika komentar tidak ditemukan
var ErrCommentNotFound = errors.New("comment not found")

// Komentar mewakili entitas komentar
type Comment struct {
    ID        uint   `gorm:"primary_key"`
    UserID    uint   `gorm:"not null"`
    PhotoID   uint   `gorm:"not null"`
    Message   string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Simpan komentar baru ke dalam database
func (c *Comment) Save() error {
    if err := DB.Create(c).Error; err != nil {
        return err
    }
    return nil
}

// GetCommentByID mengambil komentar dengan ID-nya
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

// UpdateCommentByID memperbarui komentar berdasarkan ID-nya
func UpdateCommentByID(id uint, updatedComment *Comment) error {
    if err := DB.Model(&Comment{}).Where("id = ?", id).Updates(updatedComment).Error; err != nil {
        return err
    }
    return nil
}

// DeleteCommentByID menghapus komentar berdasarkan ID-nya
func DeleteCommentByID(id uint) error {
    if err := DB.Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
        return err
    }
    return nil
}
