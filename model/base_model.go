package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id uuid.UUID `gorm:"type:uuid"`
}

func (t *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	t.Id = uuid.New()
	return
}
