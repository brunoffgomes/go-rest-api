package postgres

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"rest-api/internal/domain/entities"
)

type studentRepository struct {
	db *gorm.DB
}

func (r *studentRepository) FindById(id uint) (*entities.Student, error) {
	var student entities.Student

	result := r.db.First(&student, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}
	return &student, nil
}

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) Create(student *entities.Student) error {
	if err := r.db.Create(student).Error; err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}
	return nil
}
