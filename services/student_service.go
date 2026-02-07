package services

import (
	"errors"

	"go-api-gin-lab/models"
	"go-api-gin-lab/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func validateStudent(s models.Student, isCreate bool) error {
	if isCreate && s.Id == "" {
		return errors.New("id is required")
	}
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0.0 and 4.0")
	}
	return nil
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	if err := validateStudent(student, true); err != nil {
		return err
	}
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) error {
	if err := validateStudent(student, false); err != nil {
		return err
	}
	return s.Repo.Update(id, student)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}
