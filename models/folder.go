package models

import (
	"time"
)

type Folder struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:100;not null" json:"name"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Notes     []Note    `json:"-"`
}
