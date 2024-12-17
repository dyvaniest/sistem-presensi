package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Mahasiswa struct {
	ID       int    `gorm:"primaryKey" json:"id_mahasiswa"`
	Nama     string `json:"nama" gorm:"type:varchar(255);not null"`
	UidCard  string `json:"uid_card" gorm:"type:varchar(255);unique;not null"`
	Npm      string `json:"npm" gorm:"type:varchar(255);unique;not null"`
	Prodi    string `json:"prodi" gorm:"type:varchar(255);not null"`
	Jurusan  string `json:"jurusan" gorm:"type:varchar(255);not null"`
	Fakultas string `json:"fakultas" gorm:"type:varchar(255);not null"`
}

type Dosen struct {
	ID       int    `gorm:"primaryKey" json:"id_dosen"`
	Nama     string `json:"nama" gorm:"type:varchar(255);not null"`
	Nidn     string `json:"nidn" gorm:"type:varchar(255);unique;not null"`
	Prodi    string `json:"prodi" gorm:"type:varchar(255);not null"`
	Jurusan  string `json:"jurusan" gorm:"type:varchar(255);not null"`
	Fakultas string `json:"fakultas" gorm:"type:varchar(255);not null"`
	UserID   int    `json:"user_id" gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID;references:ID" json:"user_info"`
}

type MataKuliah struct {
	ID             int          `gorm:"primaryKey" json:"id_mk"`
	NamaMK         string       `json:"nama_mk" gorm:"type:varchar(255);not null"`
	KodeMK         string       `json:"kode_mk" gorm:"type:varchar(255);unique;not null"`
	Sks            int          `json:"sks" gorm:"not null"`
	Kelas          string       `json:"kelas" gorm:"type:varchar(10);not null"`
	JadwalKuliahID int          `json:"id_jadwal" gorm:"not null"`
	JadwalKuliah   JadwalKuliah `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:JadwalKuliahID;references:ID"`
}

type Pertemuan struct {
	ID               int        `gorm:"primaryKey" json:"id_pertemuan"`
	MataKuliahID     int        `json:"id_mk" gorm:"not null"`
	TopikPerkuliahan string     `json:"topik_perkuliahan" gorm:"type:text;not null"`
	Tanggal          time.Time  `json:"tanggal" gorm:"type:date;not null"`
	MataKuliah       MataKuliah `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MataKuliahID;references:ID"`
}

type JadwalKuliah struct {
	ID           int       `gorm:"primaryKey" json:"id_jadwal"`
	Periode      string    `json:"periode" gorm:"type:varchar(255);not null"`
	Ruang        string    `json:"ruang" gorm:"type:varchar(255);not null"`
	Hari         string    `json:"hari" gorm:"type:varchar(255);not null"`
	JamMulai     time.Time `json:"jam_mulai" gorm:"type:timestamp;not null"`
	JamSelesai   time.Time `json:"jam_selesai" gorm:"type:timestamp;not null"`
	StatusJadwal string    `json:"status_jadwal" gorm:"type:varchar(255)"`
}

type Presensi struct {
	ID              int       `gorm:"primaryKey" json:"id_presensi"`
	PertemuanID     int       `json:"id_pertemuan" gorm:"not null"`
	MahasiswaID     int       `json:"id_mahasiswa" gorm:"not null"`
	StatusKehadiran string    `json:"status_kehadiran" gorm:"type:varchar(255);not null"`
	WaktuMasuk      time.Time `json:"waktu_masuk" gorm:"not null"`
	Pertemuan       Pertemuan `gorm:"foreignKey:PertemuanID;references:ID"`
	Mahasiswa       Mahasiswa `gorm:"foreignKey:MahasiswaID;references:ID"`
}

type MahasiswaRekap struct {
	ID               int            `gorm:"primaryKey" json:"id_mahasiswa_rekap"`
	PertemuanRekapID int            `json:"id_pertemuan_rekap" gorm:"not null"`
	MahasiswaID      int            `json:"id_mahasiswa" gorm:"not null"`
	Status           string         `json:"status" gorm:"type:varchar(10);not null"`
	WaktuMasuk       *time.Time     `json:"waktu_masuk,omitempty"`
	PertemuanRekap   PertemuanRekap `gorm:"foreignKey:PertemuanRekapID;references:ID"`
}

type PertemuanRekap struct {
	ID              int           `gorm:"primaryKey" json:"id_pertemuan_rekap"`
	RekapPresensiID int           `json:"rekap_presensi_id" gorm:"not null"`
	Topik           string        `json:"topik" gorm:"type:varchar(255);not null"`
	RekapPresensi   RekapPresensi `gorm:"foreignKey:RekapPresensiID;references:ID"`
}

type RekapPresensi struct {
	ID            int              `gorm:"primaryKey" json:"id_rekap_presensi"`
	Periode       string           `json:"periode" gorm:"type:varchar(255);not null"`
	MataKuliahID  int              `json:"id_mk" gorm:"not null"`
	Kelas         string           `json:"kelas" gorm:"type:varchar(10);not null"`
	PertemuanList []PertemuanRekap `gorm:"foreignKey:RekapPresensiID" json:"pertemuan_list"`
	MataKuliah    MataKuliah       `gorm:"foreignKey:MataKuliahID;references:ID"`
}

type Session struct {
	ID       int       `gorm:"primaryKey" json:"id_session"`
	Token    string    `json:"token" gorm:"type:varchar(255);not null"`
	Username string    `json:"username" gorm:"type:varchar(255);not null"`
	Expiry   time.Time `json:"expiry" gorm:"type:timestamp;not null"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
