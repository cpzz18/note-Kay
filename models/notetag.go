package models

import "time"

type NoteTag struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    NoteID    uint      `json:"note_id"`
    TagID     uint      `json:"tag_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
