package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PertemuanAPI interface {
	AddPertemuan(c *gin.Context)
	GetPertemuanByID(c *gin.Context)
	GetAllPertemuan(c *gin.Context)
	UpdatePertemuan(c *gin.Context)
	DeletePertemuanByID(c *gin.Context)
}

type pertemuanAPI struct {
	pertemuanService services.PertemuanService
}

// new presensi api
func NewPertemuanAPI(pertemuanRepo services.PertemuanService) *pertemuanAPI {
	return &pertemuanAPI{pertemuanRepo}
}

func (p *pertemuanAPI) AddPertemuan(c *gin.Context) {
	var pertemuan models.Pertemuan
	if err := c.ShouldBindJSON(&pertemuan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recordPertemuan := models.Pertemuan{
		ID:               pertemuan.ID,
		TopikPerkuliahan: pertemuan.TopikPerkuliahan,
		Tanggal:          pertemuan.Tanggal,
		MataKuliahID:     pertemuan.MataKuliah.ID,
		MataKuliah: models.MataKuliah{
			ID:             pertemuan.MataKuliah.ID,
			NamaMK:         pertemuan.MataKuliah.NamaMK,
			KodeMK:         pertemuan.MataKuliah.KodeMK,
			Sks:            pertemuan.MataKuliah.Sks,
			Kelas:          pertemuan.MataKuliah.Kelas,
			JadwalKuliahID: pertemuan.MataKuliah.JadwalKuliahID,
			JadwalKuliah: models.JadwalKuliah{
				ID:           pertemuan.MataKuliah.JadwalKuliah.ID,
				Periode:      pertemuan.MataKuliah.JadwalKuliah.Periode,
				Ruang:        pertemuan.MataKuliah.JadwalKuliah.Ruang,
				Hari:         pertemuan.MataKuliah.JadwalKuliah.Hari,
				JamMulai:     pertemuan.MataKuliah.JadwalKuliah.JamMulai,
				JamSelesai:   pertemuan.MataKuliah.JadwalKuliah.JamSelesai,
				StatusJadwal: pertemuan.MataKuliah.JadwalKuliah.StatusJadwal,
			},
		},
	}

	if err := p.pertemuanService.AddPertemuan(&recordPertemuan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pertemuan created successfully"})
}

func (p *pertemuanAPI) GetPertemuanByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	pertemuan, err := p.pertemuanService.GetPertemuanByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pertemuan)
}

func (p *pertemuanAPI) GetAllPertemuan(c *gin.Context) {
	pertemuan, err := p.pertemuanService.GetAllPertemuan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pertemuan)
}

func (p *pertemuanAPI) UpdatePertemuan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var pertemuan models.Pertemuan
	if err := c.ShouldBindJSON(&pertemuan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.pertemuanService.UpdatePertemuan(id, pertemuan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pertemuan updated successfully"})
}

func (p *pertemuanAPI) DeletePertemuanByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := p.pertemuanService.DeletePertemuanByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pertemuan deleted successfully"})
}
