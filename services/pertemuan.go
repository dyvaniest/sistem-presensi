package services

import (
	"errors"
	"sistem-presensi/models"
	"sistem-presensi/repository"
)

type PertemuanService interface {
	AddPertemuan(pertemuan *models.Pertemuan) error
	GetPertemuanByID(id int) (*models.Pertemuan, error)
	GetAllPertemuan() ([]models.Pertemuan, error)
	UpdatePertemuan(id int, pertemuan models.Pertemuan) error
	DeletePertemuanByID(id int) error
}

type pertemuanService struct {
	repo repository.PertemuanRepository
}

func NewPertemuanService(repo repository.PertemuanRepository) PertemuanService {
	return &pertemuanService{repo}
}

// Membuat mahasiswa baru
func (p *pertemuanService) AddPertemuan(pertemuan *models.Pertemuan) error {
	// Validasi data jika diperlukan
	if pertemuan == nil {
		return errors.New("invalid pertemuan data")
	}

	// Memanggil repository untuk membuat data
	if err := p.repo.AddPertemuan(pertemuan); err != nil {
		return errors.New("failed to insert presensi")
	}
	return nil
}

func (p *pertemuanService) GetPertemuanByID(id int) (*models.Pertemuan, error) {
	pertemuan, err := p.repo.GetPertemuanByID(id)
	if err != nil {
		return nil, errors.New("presensi not found")
	}
	return &pertemuan, nil
}

func (p *pertemuanService) GetAllPertemuan() ([]models.Pertemuan, error) {
	pertemuan, err := p.repo.GetAllPertemuan()
	if err != nil {
		return nil, errors.New("failed to fetch pertemuan data")
	}
	return pertemuan, nil
}

func (p *pertemuanService) UpdatePertemuan(id int, pertemuan models.Pertemuan) error {
	// Validasi data jika diperlukan
	if pertemuan.ID == 0 {
		return errors.New("invalid pertemuan ID")
	}

	// Lakukan update melalui repository
	if err := p.repo.UpdatePertemuan(id, pertemuan); err != nil {
		return errors.New("failed to update presensi")
	}
	return nil
}

func (p *pertemuanService) DeletePertemuanByID(id int) error {
	// Validasi ID
	if id <= 0 {
		return errors.New("invalid pertemuan ID")
	}

	if err := p.repo.DeletePertemuan(id); err != nil {
		return errors.New("failed to delete presensi")
	}
	return nil
}
