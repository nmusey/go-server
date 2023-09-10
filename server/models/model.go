package models

import (
    "gorm.io/gorm"
)

type ModelInterface interface {
}

type Model struct {
    gorm.Model
    UUID string `gorm:"uniqueIndex;not null;type:uuid;default:gen_random_uuid()"`
}
