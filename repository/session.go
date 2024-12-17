package repository

import (
	"errors"
	"sistem-presensi/models"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session models.Session) error
	DeleteSession(token string) error
	UpdateSessions(session models.Session) error
	SessionAvailUsername(username string) (models.Session, error)
	SessionAvailToken(token string) (models.Session, error)
	TokenValidity(token string) (models.Session, error)
	TokenExpired(session models.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db: db}
}

func (r *sessionsRepo) AddSessions(session models.Session) error {
	result := r.db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *sessionsRepo) DeleteSession(token string) error {
	result := r.db.Where("token = ?", token).Delete(&models.Session{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}
	return nil
}

func (r *sessionsRepo) UpdateSessions(session models.Session) error {
	result := r.db.Model(&models.Session{}).Where("id = ?", session.ID).Updates(session)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}
	return nil
}

func (r *sessionsRepo) SessionAvailUsername(username string) (models.Session, error) {
	var session models.Session
	result := r.db.Where("username = ?", username).First(&session)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return session, errors.New("session not found")
		}
		return session, result.Error
	}
	return session, nil
}

func (r *sessionsRepo) SessionAvailToken(token string) (models.Session, error) {
	var session models.Session
	result := r.db.Where("token = ?", token).First(&session)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return session, errors.New("session not found")
		}
		return session, result.Error
	}
	return session, nil
}

func (r *sessionsRepo) TokenValidity(token string) (models.Session, error) {
	session, err := r.SessionAvailToken(token)
	if err != nil {
		return session, err
	}

	if r.TokenExpired(session) {
		return session, errors.New("token has expired")
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session models.Session) bool {
	return session.Expiry.Before(time.Now())
}
