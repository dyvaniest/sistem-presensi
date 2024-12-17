package services

import (
	"sistem-presensi/models"
	repo "sistem-presensi/repository"
)

type SessionService interface {
	GetSessionByEmail(email string) (models.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (s *sessionService) GetSessionByEmail(username string) (models.Session, error) {
	return s.sessionRepo.SessionAvailUsername(username)
	// return model.Session{}, nil // TODO: replace this
}
