package repository

import (
	"errors"
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type PertemuanRepository interface {
	AddPertemuan(pertemuan *models.Pertemuan) error
	GetAllPertemuan() ([]models.Pertemuan, error)
	GetPertemuanByID(id int) (models.Pertemuan, error)
	UpdatePertemuan(id int, pertemuan models.Pertemuan) error
	DeletePertemuan(id int) error
}

type pertemuanRepository struct {
	db *gorm.DB
}

func NewPertemuanRepo(db *gorm.DB) *pertemuanRepository {
	return &pertemuanRepository{db}
}

func (r *pertemuanRepository) AddPertemuan(pertemuan *models.Pertemuan) error {
	if err := r.db.Preload("MataKuliah").Preload("MataKuliah.JadwalKuliah").Create(pertemuan).Error; err != nil {
		return err
	}
	return nil
}

// Mendapatkan semua data pertemuan
func (r *pertemuanRepository) GetAllPertemuan() ([]models.Pertemuan, error) {
	var pertemuan []models.Pertemuan
	if err := r.db.Preload("MataKuliah").Preload("MataKuliah.JadwalKuliah").Find(&pertemuan).Error; err != nil {
		return nil, err
	}
	return pertemuan, nil
}

func (p *pertemuanRepository) GetPertemuanByID(id int) (models.Pertemuan, error) {
	var pertemuan models.Pertemuan
	result := p.db.Preload("MataKuliah").Preload("MataKuliah.JadwalKuliah").First(&pertemuan, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return pertemuan, errors.New("pertemuan not found")
		}
		return pertemuan, result.Error
	}
	return pertemuan, nil
}

func (p *pertemuanRepository) UpdatePertemuan(id int, pertemuan models.Pertemuan) error {
	if err := p.db.Model(&models.Presensi{}).Where("id = ?", id).Updates(pertemuan).Error; err != nil {
		return err
	}
	return nil
}

func (p *pertemuanRepository) DeletePertemuan(id int) error {
	result := p.db.Where("id = ?", id).Delete(&models.Pertemuan{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("pertemuan not found")
	}
	return nil
}
