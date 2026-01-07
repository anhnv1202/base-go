package po

import "github.com/google/uuid"

type User struct {
	UUID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username   string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	Email      string    `gorm:"type:varchar(255);not null"`
	Role       string    `gorm:"type:varchar(255);not null"`
	IsVerified bool      `gorm:"type:boolean;not null"`
}
