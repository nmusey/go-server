package models

type User struct {
    Model

    Email string `gorm:"uniqueIndex;not null"`
}
