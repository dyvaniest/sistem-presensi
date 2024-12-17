package services

import (
	"errors"
	"sistem-presensi/models"
	"sistem-presensi/repository"
)

type MataKuliahService interface {
	GetMataKuliahByID(id int) (*models.MataKuliah, error)
	UpdateMataKuliah(id int, mataKuliah models.MataKuliah) error
	DeleteMataKuliahByID(id int) error
}

type mataKuliahService struct {
	repo repository.MataKuliahRepository
}

// Constructor untuk membuat instance MataKuliahService
func NewMataKuliahService(repo repository.MataKuliahRepository) MataKuliahService {
	return &mataKuliahService{repo}
}

// Mendapatkan mata kuliah berdasarkan ID
func (mk *mataKuliahService) GetMataKuliahByID(id int) (*models.MataKuliah, error) {
	if id <= 0 {
		return nil, errors.New("invalid mata kuliah ID")
	}

	mataKuliah, err := mk.repo.GetMataKuliahByID(id)
	if err != nil {
		return nil, errors.New("mata kuliah not found")
	}
	return mataKuliah, nil
}

// Memperbarui data mata kuliah
func (mk *mataKuliahService) UpdateMataKuliah(id int, mataKuliah models.MataKuliah) error {
	if id <= 0 {
		return errors.New("invalid mata kuliah ID")
	}

	// Memastikan mata kuliah ada sebelum memperbarui
	_, err := mk.repo.GetMataKuliahByID(id)
	if err != nil {
		return errors.New("mata kuliah not found")
	}

	// Memanggil repository untuk memperbarui data
	if err := mk.repo.UpdateMataKuliah(id, mataKuliah); err != nil {
		return errors.New("failed to update mata kuliah")
	}
	return nil
}

// Menghapus mata kuliah berdasarkan ID
func (mk *mataKuliahService) DeleteMataKuliahByID(id int) error {
	if id <= 0 {
		return errors.New("invalid mata kuliah ID")
	}

	// Memastikan mata kuliah ada sebelum menghapus
	_, err := mk.repo.GetMataKuliahByID(id)
	if err != nil {
		return errors.New("mata kuliah not found")
	}

	// Memanggil repository untuk menghapus data
	if err := mk.repo.DeleteMataKuliahByID(id); err != nil {
		return errors.New("failed to delete mata kuliah")
	}
	return nil
}
