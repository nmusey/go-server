package utils

import (
    "gorm.io/gorm"

    "server/data"
    "server/models"
    "server/services"
)

type DependencyInjector struct {
    DB *gorm.DB

    UserRepository *data.Repository[models.User]
    UserService *services.UserService
}

func NewDIContainer() *DependencyInjector {
    db, err := data.ConnectToDatabase()
    if err != nil {
        panic(err)
    }

    userRepository := data.NewRepository[models.User](db)
    userService := services.NewUserService(userRepository)

    return &DependencyInjector{
        DB: db,
        UserRepository: userRepository,
        UserService: userService,
    }
}
