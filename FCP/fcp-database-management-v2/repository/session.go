package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	return model.Session{}, nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	return model.Session{}, nil // TODO: replace this
}
