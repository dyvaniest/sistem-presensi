package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JadwalAPI interface {
	AddJadwal(c *gin.Context)
	GetJadwalByID(c *gin.Context)
	UpdateJadwal(c *gin.Context)
	DeleteJadwalByID(c *gin.Context)
}

type jadwalAPI struct {
	jadwalService services.JadwalService
}

func NewJadwalAPI(jadwalRepo services.JadwalService) *jadwalAPI {
	return &jadwalAPI{jadwalRepo}
}

func (j *jadwalAPI) AddJadwal(c *gin.Context) {
	var jadwal models.JadwalKuliah
	if err := c.ShouldBindJSON(&jadwal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := j.jadwalService.AddJadwal(&jadwal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Jadwal kuliah created successfully"})
}

func (j *jadwalAPI) GetJadwalByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jadwal, err := j.jadwalService.GetJadwalKuliahByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jadwal)
}

func (j *jadwalAPI) UpdateJadwal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jadwal models.JadwalKuliah
	if err := c.ShouldBindJSON(&jadwal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := j.jadwalService.UpdateJadwalKuliah(id, jadwal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jadwal updated successfully"})
}

func (j *jadwalAPI) DeleteJadwalByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := j.jadwalService.DeleteJadwalKuliahByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jadwal deleted successfully"})
}
