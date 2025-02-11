package entities

import "gorm.io/gorm"

type Discipline struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string
}
