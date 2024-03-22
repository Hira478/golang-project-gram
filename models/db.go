// db.go

package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
    var err error
    // Replace the connection string placeholders with actual connection details
    connectionString := fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PASSWORD"),
    )
    DB, err = gorm.Open("postgres", connectionString)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    // Enable logging SQL statements (optional)
    DB.LogMode(true)

    // Migrate the database schemas
    migrateDatabase()
}

func migrateDatabase() {
    // Auto migrate database schemas
    err := DB.AutoMigrate(&User{}, &Photo{}, &Comment{}, &SocialMedia{}).Error
    if err != nil {
        log.Fatal("Error migrating database schemas:", err)
    }
}
