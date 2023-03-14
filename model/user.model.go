package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirsName  string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
