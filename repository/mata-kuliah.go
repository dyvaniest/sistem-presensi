package repository

import (
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	GetMataKuliahByID(id int) (*models.MataKuliah, error)
	UpdateMataKuliah(id int, mataKuliah models.MataKuliah) error
	DeleteMataKuliahByID(id int) error
}

type mataKuliahRepository struct {
	db *gorm.DB
}

func NewMataKuliahRepository(db *gorm.DB) *mataKuliahRepository {
	return &mataKuliahRepository{db: db}
}

func (r *mataKuliahRepository) GetMataKuliahByID(id int) (*models.MataKuliah, error) {
	var matkul models.MataKuliah
	if err := r.db.Preload("JadwalKuliah").First(&matkul, id).Error; err != nil {
		return nil, err
	}
	return &matkul, nil
}

func (r *mataKuliahRepository) UpdateMataKuliah(id int, mataKuliah models.MataKuliah) error {
	if err := r.db.Preload("JadwalKuliah").Model(&models.MataKuliah{}).Where("id = ?", id).Updates(mataKuliah).Error; err != nil {
		return err
	}
	return nil
}

func (r *mataKuliahRepository) DeleteMataKuliahByID(id int) error {
	if err := r.db.Preload("JadwalKuliah").Delete(&models.MataKuliah{}, id).Error; err != nil {
		return err
	}
	return nil
}
