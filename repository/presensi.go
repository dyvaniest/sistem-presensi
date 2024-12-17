package repository

import (
	"errors"
	"sistem-presensi/models"

	"gorm.io/gorm"
)

type PresensiRepository interface {
	FindMahasiswaByUID(uid string) (*models.Mahasiswa, error)
	FindPertemuanByID(id int) (*models.Pertemuan, error)
	InsertPresensi(presensi *models.Presensi) error
	GetPresensiByID(id int) (*models.Presensi, error)
	GetAllPresensi() ([]models.Presensi, error)
	UpdatePresensi(id int, presensi models.Presensi) error
	DeletePresensiByID(id int) error
}

type presensiRepository struct {
	db *gorm.DB
}

func NewPresensiRepository(db *gorm.DB) *presensiRepository {
	return &presensiRepository{db: db}
}

func (r *presensiRepository) FindMahasiswaByUID(uid string) (*models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	if err := r.db.Where("uid_card = ?", uid).First(&mahasiswa).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("mahasiswa tidak ditemukan")
		}
		return nil, err
	}
	return &mahasiswa, nil
}

func (r *presensiRepository) FindPertemuanByID(id int) (*models.Pertemuan, error) {
	var pertemuan models.Pertemuan
	if err := r.db.Preload("MataKuliah").
		Preload("MataKuliah.JadwalKuliah").
		First(&pertemuan, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("pertemuan tidak ditemukan")
		}
		return nil, err
	}
	return &pertemuan, nil
}

func (r *presensiRepository) InsertPresensi(presensi *models.Presensi) error {
	if err := r.db.Create(presensi).Error; err != nil {
		return err
	}
	return nil
}

func (r *presensiRepository) GetPresensiByID(id int) (*models.Presensi, error) {
	var presensi models.Presensi
	if err := r.db.
		Preload("Pertemuan").
		Preload("Pertemuan.MataKuliah").
		Preload("Pertemuan.MataKuliah.JadwalKuliah").
		Preload("Mahasiswa").
		First(&presensi, id).Error; err != nil {
		return nil, err
	}
	return &presensi, nil
}

func (r *presensiRepository) GetAllPresensi() ([]models.Presensi, error) {
	var presensi []models.Presensi
	if err := r.db.
		Preload("Pertemuan").
		Preload("Pertemuan.MataKuliah").
		Preload("Pertemuan.MataKuliah.JadwalKuliah").
		Preload("Mahasiswa").
		Find(&presensi).Error; err != nil {
		return nil, err
	}
	return presensi, nil
}

func (r *presensiRepository) UpdatePresensi(id int, presensi models.Presensi) error {
	if err := r.db.Model(&models.Presensi{}).Where("id = ?", id).Updates(presensi).Error; err != nil {
		return err
	}
	return nil
}

func (r *presensiRepository) DeletePresensiByID(id int) error {
	if err := r.db.Delete(&models.Presensi{}, id).Error; err != nil {
		return err
	}
	return nil
}
