package services

import (
	"errors"
	"fmt"
	"sistem-presensi/models"
	"sistem-presensi/repository"
	"time"
)

type PresensiService interface {
	RecordPresensi(uid string, pertemuanID int) (string, error)
	InsertPresensi(presensi *models.Presensi) error
	GetPresensiByID(id int) (*models.Presensi, error)
	GetAllPresensi() ([]models.Presensi, error)
	UpdatePresensi(id int, presensi models.Presensi) error
	DeletePresensiByID(id int) error
}

type presensiService struct {
	repo repository.PresensiRepository
}

// Constructor untuk membuat instance PresensiService
func NewPresensiService(repo repository.PresensiRepository) PresensiService {
	return &presensiService{repo}
}

func (s *presensiService) RecordPresensi(uid string, pertemuanID int) (string, error) {
	// Cek mahasiswa berdasarkan UID
	mahasiswa, err := s.repo.FindMahasiswaByUID(uid)
	if err != nil {
		return "", fmt.Errorf("gagal menemukan mahasiswa dengan UID: %s, error: %w", uid, err)
	}

	// Cek pertemuan berdasarkan ID
	pertemuan, err := s.repo.FindPertemuanByID(pertemuanID)
	if err != nil {
		return "", fmt.Errorf("gagal menemukan pertemuan dengan ID: %d, error: %w", pertemuanID, err)
	}

	// Validasi MataKuliah
	if pertemuan.MataKuliah.ID == 0 {
		return "", errors.New("data mata kuliah tidak valid")
	}

	// Validasi JadwalKuliah
	if pertemuan.MataKuliah.JadwalKuliah.ID == 0 {
		return "", errors.New("data jadwal kuliah tidak valid")
	}

	// Validasi waktu presensi
	currentTime := time.Now()
	jamMulai := pertemuan.MataKuliah.JadwalKuliah.JamMulai
	jamSelesai := pertemuan.MataKuliah.JadwalKuliah.JamSelesai

	if currentTime.Before(jamMulai) {
		return "", fmt.Errorf("waktu presensi belum dimulai (mulai: %s)", jamMulai.Format("15:04:05"))
	}

	if currentTime.After(jamSelesai) {
		return "", fmt.Errorf("waktu presensi sudah berakhir (berakhir: %s)", jamSelesai.Format("15:04:05"))
	}

	// Rekam presensi
	presensi := &models.Presensi{
		PertemuanID:     pertemuan.ID,
		MahasiswaID:     mahasiswa.ID,
		StatusKehadiran: "Hadir",
		WaktuMasuk:      currentTime,
	}

	if err := s.repo.InsertPresensi(presensi); err != nil {
		return "", fmt.Errorf("gagal merekam presensi: %w", err)
	}

	return "Presensi berhasil direkam", nil
}

// Membuat mahasiswa baru
func (p *presensiService) InsertPresensi(presensi *models.Presensi) error {
	// Validasi data jika diperlukan
	if presensi == nil {
		return errors.New("invalid presensi data")
	}

	// Memanggil repository untuk membuat data
	if err := p.repo.InsertPresensi(presensi); err != nil {
		return errors.New("failed to insert presensi")
	}
	return nil
}

// Mendapatkan presensi berdasarkan ID
func (p *presensiService) GetPresensiByID(id int) (*models.Presensi, error) {
	presensi, err := p.repo.GetPresensiByID(id)
	if err != nil {
		return nil, errors.New("presensi not found")
	}
	return presensi, nil
}

// Mendapatkan semua data presensi
func (p *presensiService) GetAllPresensi() ([]models.Presensi, error) {
	presensi, err := p.repo.GetAllPresensi()
	if err != nil {
		return nil, errors.New("failed to fetch presensi data")
	}
	return presensi, nil
}

// Memperbarui data presensi
func (p *presensiService) UpdatePresensi(id int, presensi models.Presensi) error {
	// Validasi data jika diperlukan
	if presensi.ID == 0 {
		return errors.New("invalid presensi ID")
	}

	// Lakukan update melalui repository
	if err := p.repo.UpdatePresensi(id, presensi); err != nil {
		return errors.New("failed to update presensi")
	}
	return nil
}

// Menghapus presensi berdasarkan ID
func (p *presensiService) DeletePresensiByID(id int) error {
	// Validasi ID
	if id <= 0 {
		return errors.New("invalid presensi ID")
	}

	if err := p.repo.DeletePresensiByID(id); err != nil {
		return errors.New("failed to delete presensi")
	}
	return nil
}
