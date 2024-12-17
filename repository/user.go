package repository

import (
	"errors"
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUserByUsername(username string) error
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, result.Error
	}
	return user, nil
}

// fungsi UpdateUser
func (r *userRepository) UpdateUser(user models.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// fungsi DeleteUserByUsername
func (r *userRepository) DeleteUserByUsername(username string) error {
	result := r.db.Where("username = ?", username).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// fungsi CreateUser
func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
