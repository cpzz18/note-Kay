package models

import (
	"time"
)

type Note struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	FolderID  uint      `json:"folder_id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	
	Folder   *Folder   `json:"folder,omitempty"`
	NoteTags []NoteTag `json:"note_tags,omitempty"`
}
