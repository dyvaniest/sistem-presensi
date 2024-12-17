package services

import (
	"errors"
	"sistem-presensi/models"
	"sistem-presensi/repository"
)

type JadwalService interface {
	AddJadwal(jadwal *models.JadwalKuliah) error
	GetJadwalKuliahByID(id int) (*models.JadwalKuliah, error)
	UpdateJadwalKuliah(id int, jadwal models.JadwalKuliah) error
	DeleteJadwalKuliahByID(id int) error
}

type jadwalService struct {
	repo repository.JadwalRespository
}

// Constructor untuk membuat instance JadwalService
func NewJadwalService(repo repository.JadwalRespository) JadwalService {
	return &jadwalService{repo}
}

// Membuat jadwal baru
func (j *jadwalService) AddJadwal(jadwal *models.JadwalKuliah) error {
	// Validasi data jika diperlukan
	if jadwal == nil {
		return errors.New("invalid jadwal data")
	}

	// Memanggil repository untuk membuat data
	if err := j.repo.AddJadwal(jadwal); err != nil {
		return errors.New("failed to create jadwal")
	}
	return nil
}

// Mendapatkan jadwal kuliah berdasarkan ID
func (j *jadwalService) GetJadwalKuliahByID(id int) (*models.JadwalKuliah, error) {
	if id <= 0 {
		return nil, errors.New("invalid jadwal kuliah ID")
	}

	jadwal, err := j.repo.GetJadwalKuliahByID(id)
	if err != nil {
		return nil, errors.New("jadwal kuliah not found")
	}
	return jadwal, nil
}

// Memperbarui data jadwal kuliah
func (j *jadwalService) UpdateJadwalKuliah(id int, jadwal models.JadwalKuliah) error {
	if id <= 0 {
		return errors.New("invalid jadwal kuliah ID")
	}

	// Memastikan jadwal kuliah ada sebelum memperbarui
	_, err := j.repo.GetJadwalKuliahByID(id)
	if err != nil {
		return errors.New("jadwal kuliah not found")
	}

	// Memanggil repository untuk memperbarui data
	if err := j.repo.UpdateJadwalKuliah(id, jadwal); err != nil {
		return errors.New("failed to update jadwal kuliah")
	}
	return nil
}

// Menghapus jadwal kuliah berdasarkan ID
func (j *jadwalService) DeleteJadwalKuliahByID(id int) error {
	if id <= 0 {
		return errors.New("invalid jadwal kuliah ID")
	}

	// Memastikan jadwal kuliah ada sebelum menghapus
	_, err := j.repo.GetJadwalKuliahByID(id)
	if err != nil {
		return errors.New("jadwal kuliah not found")
	}

	// Memanggil repository untuk menghapus data
	if err := j.repo.DeleteJadwalKuliahByID(id); err != nil {
		return errors.New("failed to delete jadwal kuliah")
	}
	return nil
}
