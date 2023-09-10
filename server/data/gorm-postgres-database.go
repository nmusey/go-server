package data

import (
    "fmt"
    "os"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func ConnectToDatabase() (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", 
        os.Getenv("DB_HOST"), 
        os.Getenv("DB_USER"), 
        os.Getenv("DB_PASS"), 
        os.Getenv("DB_NAME"), 
        os.Getenv("DB_PORT"),
    )

    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
