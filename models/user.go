// user.go

package models

import (
	"time"

	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
    ID        uint   `gorm:"primary_key"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    Age       int    `gorm:"not null;check:age > 8"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (u *User) Save() error {
    if err := DB.Create(&u).Error; err != nil {
        return err
    }
    return nil
}

// Update updates an existing user record in the database
func (u *User) Update(updatedUser *User) error {
    if err := DB.Model(u).Updates(updatedUser).Error; err != nil {
        return err
    }
    return nil
}

// Delete deletes a user record from the database
func (u *User) Delete() error {
    if err := DB.Delete(&u).Error; err != nil {
        return err
    }
    return nil
}

// ErrUserNotFound is the error returned when a user is not found
var ErrUserNotFound = errors.New("user not found")

// GetUserByEmail retrieves a user by their email address
func GetUserByEmail(email string) (*User, error) {
    var user User
    if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

// GetUserByID retrieves a user by their ID from the database
func GetUserByID(userID uint) (*User, error) {
    var user User
    if err := DB.First(&user, userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}