package services

import (
	"errors"
	"sistem-presensi/models"
	"sistem-presensi/repository"
)

type DosenService interface {
	AddDosen(dosen models.Dosen) error
	GetDosenByID(id int) (*models.Dosen, error)
	GetAllDosen() ([]models.Dosen, error)
}

type dosenService struct {
	repo repository.DosenRepository
}

// Constructor untuk DosenService
func NewDosenService(repo repository.DosenRepository) DosenService {
	return &dosenService{repo}
}

// Menambahkan data dosen
func (d *dosenService) AddDosen(dosen models.Dosen) error {
	if dosen.Nama == "" {
		return errors.New("nama dosen tidak boleh kosong")
	}
	return d.repo.AddDosen(dosen)
}

// Mendapatkan data dosen berdasarkan ID
func (d *dosenService) GetDosenByID(id int) (*models.Dosen, error) {
	if id <= 0 {
		return nil, errors.New("invalid dosen ID")
	}

	dosen, err := d.repo.GetDosenByID(id)
	if err != nil {
		return nil, errors.New("dosen not found")
	}
	return dosen, nil
}

// Mendapatkan semua data dosen
func (d *dosenService) GetAllDosen() ([]models.Dosen, error) {
	dosen, err := d.repo.GetAllDosen()
	if err != nil {
		return nil, errors.New("failed to retrieve dosen list")
	}
	return dosen, nil
}
