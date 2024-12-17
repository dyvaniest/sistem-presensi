package repository

import (
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type DosenRepository interface {
	AddDosen(dosen models.Dosen) error
	GetDosenByID(id int) (*models.Dosen, error)
	GetAllDosen() ([]models.Dosen, error)
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepo(db *gorm.DB) *dosenRepository {
	return &dosenRepository{db}
}

func (r *dosenRepository) AddDosen(dosen models.Dosen) error {
	if err := r.db.Create(&dosen).Error; err != nil {
		return err
	}
	return nil
}

// Mendapatkan data dosen berdasarkan ID
func (r *dosenRepository) GetDosenByID(id int) (*models.Dosen, error) {
	var dosen models.Dosen
	if err := r.db.First(&dosen, id).Error; err != nil {
		return nil, err
	}
	return &dosen, nil
}

// Mendapatkan semua data dosen
func (r *dosenRepository) GetAllDosen() ([]models.Dosen, error) {
	var dosen []models.Dosen
	if err := r.db.Find(&dosen).Error; err != nil {
		return nil, err
	}
	return dosen, nil
}
