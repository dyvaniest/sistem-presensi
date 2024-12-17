package services

import (
	"errors"
	"sistem-presensi/models"
	"sistem-presensi/repository"
)

type MahasiswaService interface {
	GetAllMahasiswa() ([]models.Mahasiswa, error)
	GetMahasiswaByID(id int) (*models.Mahasiswa, error)
	CreateMahasiswa(mahasiswa *models.Mahasiswa) error
	UpdateMahasiswa(id int, mahasiswa *models.Mahasiswa) error
	DeleteMahasiswa(id int) error
}

type mahasiswaService struct {
	repo repository.MahasiswaRepository
}

// Constructor untuk membuat instance MahasiswaService
func NewMahasiswaService(repo repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{repo}
}

// Mendapatkan semua mahasiswa
func (m *mahasiswaService) GetAllMahasiswa() ([]models.Mahasiswa, error) {
	mahasiswa, err := m.repo.GetAll()
	if err != nil {
		return nil, errors.New("failed to fetch mahasiswa data")
	}
	return mahasiswa, nil
}

// Mendapatkan mahasiswa berdasarkan ID
func (m *mahasiswaService) GetMahasiswaByID(id int) (*models.Mahasiswa, error) {
	mahasiswa, err := m.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("mahasiswa not found")
	}
	return mahasiswa, nil
}

// Membuat mahasiswa baru
func (m *mahasiswaService) CreateMahasiswa(mahasiswa *models.Mahasiswa) error {
	// Validasi data jika diperlukan
	if mahasiswa == nil {
		return errors.New("invalid mahasiswa data")
	}

	// Memanggil repository untuk membuat data
	if err := m.repo.Create(mahasiswa); err != nil {
		return errors.New("failed to create mahasiswa")
	}
	return nil
}

// Memperbarui data mahasiswa
func (m *mahasiswaService) UpdateMahasiswa(id int, mahasiswa *models.Mahasiswa) error {
	// Validasi ID
	if id <= 0 {
		return errors.New("invalid mahasiswa ID")
	}

	// Memanggil repository untuk memperbarui data
	if err := m.repo.Update(id, mahasiswa); err != nil {
		return errors.New("failed to update mahasiswa")
	}
	return nil
}

// Menghapus mahasiswa berdasarkan ID
func (m *mahasiswaService) DeleteMahasiswa(id int) error {
	// Validasi ID
	if id <= 0 {
		return errors.New("invalid mahasiswa ID")
	}

	// Memanggil repository untuk menghapus data
	if err := m.repo.Delete(id); err != nil {
		return errors.New("failed to delete mahasiswa")
	}
	return nil
}
