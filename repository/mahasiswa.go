package repository

import (
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	GetAll() ([]models.Mahasiswa, error)
	GetByID(id int) (*models.Mahasiswa, error)
	Create(mahasiswa *models.Mahasiswa) error
	Update(id int, mahasiswa *models.Mahasiswa) error
	Delete(id int) error
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) *mahasiswaRepository {
	return &mahasiswaRepository{db: db}
}

func (r *mahasiswaRepository) GetAll() ([]models.Mahasiswa, error) {
	var mahasiswa []models.Mahasiswa
	if err := r.db.Find(&mahasiswa).Error; err != nil {
		return nil, err
	}
	return mahasiswa, nil
}

func (r *mahasiswaRepository) GetByID(id int) (*models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	if err := r.db.First(&mahasiswa, id).Error; err != nil {
		return nil, err
	}
	return &mahasiswa, nil
}

func (r *mahasiswaRepository) Create(mahasiswa *models.Mahasiswa) error {
	if err := r.db.Create(mahasiswa).Error; err != nil {
		return err
	}
	return nil
}

func (r *mahasiswaRepository) Update(id int, mahasiswa *models.Mahasiswa) error {
	if err := r.db.Model(&models.Mahasiswa{}).Where("id = ?", id).Updates(mahasiswa).Error; err != nil {
		return err
	}
	return nil
}

func (r *mahasiswaRepository) Delete(id int) error {
	if err := r.db.Delete(&models.Mahasiswa{}, id).Error; err != nil {
		return err
	}
	return nil
}
