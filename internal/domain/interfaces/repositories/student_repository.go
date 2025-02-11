package repositories

import "rest-api/internal/domain/entities"

type StudentRepository interface {
	Create(user *entities.Student)
	FindById(id uint) (*entities.Student, error)
}
