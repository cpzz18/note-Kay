package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	UserID uint
	Name   string
}
