package entities

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Value        float64
	StudentID    uint `gorm:"primaryKey"`
	DisciplineID uint `gorm:"primaryKey"`
}
