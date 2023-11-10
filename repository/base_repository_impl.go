package repository

import (
	"errors"

	"github.com/f-pochat/golang-base/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepositoryImpl[T any] struct {
	Db *gorm.DB
}

func NewBaseRepositoryImpl[T any](Db *gorm.DB) BaseRepository[T] {
	return &BaseRepositoryImpl[T]{Db: Db}
}

func (repo BaseRepositoryImpl[T]) Save(data T) {
	result := repo.Db.Create(&data)
	helper.ErrorPanic(result.Error)

}

func (repo BaseRepositoryImpl[T]) Update(data T) {
	result := repo.Db.Model(&data).Updates(data)
	helper.ErrorPanic(result.Error)
}

func (repo BaseRepositoryImpl[T]) Delete(id uuid.UUID) {
	var data T
	result := repo.Db.Where("id = ?", id).Delete(&data)
	helper.ErrorPanic(result.Error)
}

func (repo BaseRepositoryImpl[T]) FindById(id uuid.UUID) (T, error) {
	var data T
	result := repo.Db.Find(&data, id)
	if result != nil {
		return data, nil
	} else {
		return data, errors.New("id not found")
	}
}

func (repo BaseRepositoryImpl[T]) FindAll() []T {
	var data []T
	results := repo.Db.Find(&data)
	helper.ErrorPanic(results.Error)
	return data
}
