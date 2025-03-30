package models

import "time"

type File struct {
	ID         uint      `gorm:"primaryKey"`
	Filename   string    `gorm:"not null"`
	Size       int64     `gorm:"not null"`
	UploadedAt time.Time `gorm:"default:current_timestamp"`
}
