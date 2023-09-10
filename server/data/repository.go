package data

import (
    "gorm.io/gorm"
    
    "server/models"
)

type Repository[T models.ModelInterface] struct {
    db *gorm.DB
}

func NewRepository[T models.ModelInterface](db *gorm.DB) *Repository[T] {
    return &Repository[T] {
        db: db,
    }
}

func (r *Repository[T]) FindById(id uint) (T, error) {
    var model T
    result := r.db.First(model, id)
    return model, result.Error
}

func (r *Repository[T]) FindByUUID(uuid string) (T, error) {
    var model T
    result := r.db.Where("uuid = ?", uuid).First(model)
    return model, result.Error
}

func (r *Repository[T]) FindAll() ([]T, error) {
    models := []T{}
    result := r.db.Find(&models)
    return models, result.Error
}

func (r *Repository[T]) FindWhere(query string, args ...string) ([]T, error) {
    models := []T{}
    result := r.db.Where(query, args).Find(&models)
    return models, result.Error
}

func (r *Repository[T]) Create(model T) error {
    return r.db.Create(model).Error
}

func (r *Repository[T]) Update(model T) error {
    return r.db.Save(model).Error
}

func (r *Repository[T]) Delete(model T) error {
    return r.db.Delete(model).Error
}

