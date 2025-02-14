package student

import (
	"fmt"
	"rest-api/internal/domain/entities"
	"rest-api/internal/interfaces/repositories"
)

type FindStudentByIdInput struct {
	Id uint
}

type GetStudentByIDUseCase struct {
	studentRepo repositories.StudentRepository
}

func GetStudentByIdStudentUseCase(studentRepo repositories.StudentRepository) *GetStudentByIDUseCase {
	return &GetStudentByIDUseCase{
		studentRepo: studentRepo,
	}
}

func (uc *GetStudentByIDUseCase) Execute(id uint) (*entities.Student, error) {

	student, err := uc.studentRepo.FindById(id)

	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	if student == nil {
		return nil, fmt.Errorf("user not found with id: %d", id)
	}

	return student, nil
}
