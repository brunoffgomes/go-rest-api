package entities

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id        uint `gorm:"primaryKey"`
	Name      string
	Matricula string
	Grades    []Grade
}
