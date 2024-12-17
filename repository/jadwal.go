package repository

import (
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type JadwalRespository interface {
	AddJadwal(jadwal *models.JadwalKuliah) error
	GetJadwalKuliahByID(id int) (*models.JadwalKuliah, error)
	UpdateJadwalKuliah(id int, jadwal models.JadwalKuliah) error
	DeleteJadwalKuliahByID(id int) error
}

type jadwalRespository struct {
	db *gorm.DB
}

func NewJadwalRespository(db *gorm.DB) *jadwalRespository {
	return &jadwalRespository{db: db}
}

func (r *jadwalRespository) AddJadwal(jadwal *models.JadwalKuliah) error {
	if err := r.db.Create(jadwal).Error; err != nil {
		return err
	}
	return nil
}

func (r *jadwalRespository) GetJadwalKuliahByID(id int) (*models.JadwalKuliah, error) {
	var jadwal models.JadwalKuliah
	if err := r.db.First(&jadwal, id).Error; err != nil {
		return nil, err
	}
	return &jadwal, nil
}

func (r *jadwalRespository) UpdateJadwalKuliah(id int, jadwal models.JadwalKuliah) error {
	if err := r.db.Model(&models.JadwalKuliah{}).Where("id = ?", id).Updates(jadwal).Error; err != nil {
		return err
	}
	return nil
}

func (r *jadwalRespository) DeleteJadwalKuliahByID(id int) error {
	if err := r.db.Delete(&models.JadwalKuliah{}, id).Error; err != nil {
		return err
	}
	return nil
}
