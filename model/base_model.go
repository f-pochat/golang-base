package model

import (
	"github.com/google/uuid"
)

type BaseModel struct {
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
