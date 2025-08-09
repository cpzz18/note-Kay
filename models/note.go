package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID   uint
	FolderID *uint
	Title    string
	Content  string
	Tags     []Tag `gorm:"many2many":note_tags;`
}
