// models/word.go
package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
}
