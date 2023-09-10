package services

import (
    "server/data"
    "server/models"
)

type UserService struct {
    Service[models.User]
}

func NewUserService(repo *data.Repository[models.User]) *UserService {
    service := NewService[models.User](repo)
    return &UserService{Service: *service}
}
