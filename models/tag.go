package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Notes []Note `gorm:"many2many:note_tags;"`
}
