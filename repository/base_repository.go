package repository

import (
	"github.com/google/uuid"
)

// T must contain an id of type uuid
type BaseRepository[T any] interface {
	Save(data T)
	Update(data T)
	Delete(id uuid.UUID)
	FindById(id uuid.UUID) (data T, err error)
	FindAll() []T
}
