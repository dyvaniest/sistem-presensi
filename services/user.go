package services

import (
	"errors"
	"sistem-presensi/models"
	repo "sistem-presensi/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *models.User) (models.User, error)
	Login(user *models.User) (token *string, err error)
	GetUserByUsername(username string) (models.User, error)
}

type userService struct {
	userRepo     repo.UserRepository
	sessionsRepo repo.SessionRepository
}

func NewUserService(userRepository repo.UserRepository, sessionsRepo repo.SessionRepository) UserService {
	return &userService{userRepository, sessionsRepo}
}

func (s *userService) Register(user *models.User) (models.User, error) {
	dbUser, err := s.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return *user, err
	}

	if dbUser.Username != "" || dbUser.ID != 0 {
		return *user, errors.New("username already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (u *userService) Login(user *models.User) (token *string, err error) {
	dbUser, err := u.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	if dbUser.Username == "" || dbUser.ID == 0 {
		return nil, errors.New("user not found")
	}

	if user.Password != dbUser.Password {
		return nil, errors.New("wrong email or password")
	}

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &models.Claims{
		Username: dbUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(models.JwtKey)
	if err != nil {
		return nil, err
	}

	session := models.Session{
		Token:    tokenString,
		Username: user.Username,
		Expiry:   expirationTime,
	}

	_, err = u.sessionsRepo.SessionAvailUsername(session.Username)
	if err != nil {
		err = u.sessionsRepo.AddSessions(session)
	} else {
		err = u.sessionsRepo.UpdateSessions(session)
	}

	return &tokenString, err
}

func (u *userService) GetUserByUsername(username string) (models.User, error) {
	return u.userRepo.GetUserByUsername(username)
}
