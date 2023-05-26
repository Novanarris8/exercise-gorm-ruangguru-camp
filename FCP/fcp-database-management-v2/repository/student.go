package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	return []model.Student{}, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	return fmt.Errorf("not implement") // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	return nil, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	return nil, nil // TODO: replace this
}
