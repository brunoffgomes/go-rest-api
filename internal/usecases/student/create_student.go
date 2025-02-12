package student

import (
	"fmt"
	"rest-api/internal/domain/entities"
	"rest-api/internal/domain/interfaces/repositories"
)

type CreateStudentInput struct {
	Name      string
	Matricula string
}

type CreateStudentUseCase struct {
	studentRepo repositories.StudentRepository
}

func NewCreateStudentUseCase(studentRepo repositories.StudentRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{
		studentRepo: studentRepo,
	}
}

func (uc *CreateStudentUseCase) Execute(input *CreateStudentInput) (*entities.Student, error) {
	student := &entities.Student{
		Name:      input.Name,
		Matricula: input.Matricula,
	}

	if err := uc.studentRepo.Create(student); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return student, nil
}
