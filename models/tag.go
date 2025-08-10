package models

import "time"

type Tag struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:50;not null;unique" json:"name"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Notes     []NoteTag `json:"-"`
}