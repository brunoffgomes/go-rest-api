package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"rest-api/internal/domain/entities"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) Create(user *entities.Student) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}
	return nil
}

func (r *studentRepository) FindByID(id uint) (*entities.Student, error) {
	var user entities.Student
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("Failed to find user: %w", err)
	}
	return &user, nil
}
