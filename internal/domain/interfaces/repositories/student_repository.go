package repositories

import "rest-api/internal/domain/entities"

type StudentRepository interface {
	Create(student *entities.Student) error
	//FindById(id uint) (*entities.Student, error)
}
