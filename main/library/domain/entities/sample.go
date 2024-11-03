package domain

import "gorm.io/gorm"

type Sample struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"ID"`
	gorm.Model
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Size      float64 `json:"size"`
	IsVisible bool    `json:"is_visible"`
}
