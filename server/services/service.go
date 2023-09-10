package services

import (
    "server/data"
    "server/models"
)

type Service[T models.ModelInterface] struct {
    repo *data.Repository[T]
}

func NewService[T models.ModelInterface](repo *data.Repository[T]) *Service[T] {
    return &Service[T]{repo: repo}
}

func (s *Service[T]) FindById(id uint) (T, error) {
    return s.repo.FindById(id)
}

func (s *Service[T]) FindByUUID(uuid string) (T, error) {
    return s.repo.FindByUUID(uuid)
}

func (s *Service[T]) FindAll() ([]T, error) {
    return s.repo.FindAll()
}

func (s *Service[T]) FindWhere(query string, args ...string) ([]T, error) {
    return s.repo.FindWhere(query, args...)
}

func (s *Service[T]) Create(model T) error {
    return s.repo.Create(model)
}

func (s *Service[T]) Update(model T) error {
    return s.repo.Update(model)
}

func (s *Service[T]) Delete(model T) error {
    return s.repo.Delete(model)
}
