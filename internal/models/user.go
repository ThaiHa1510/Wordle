package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"` // In practice, store hashed passwords
	Words    []Word
}
