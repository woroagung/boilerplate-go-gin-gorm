package example

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Example struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name string    `gorm:"type:varchar(255);not null;" json:"name"`
}

// hooks
func (e *Example) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

// your functions here
