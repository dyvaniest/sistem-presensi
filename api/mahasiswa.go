package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MahasiswaAPI interface {
	GetMahasiswaByID(c *gin.Context)
	GetAllMahasiswa(c *gin.Context)
	CreateMahasiswa(c *gin.Context)
	UpdateMahasiswa(c *gin.Context)
	DeleteMahasiswaByID(c *gin.Context)
}

type mahasiswaAPI struct {
	mhsService services.MahasiswaService
}

func NewMahasiswaAPI(mhsRepo services.MahasiswaService) *mahasiswaAPI {
	return &mahasiswaAPI{mhsRepo}
}

func (m *mahasiswaAPI) GetMahasiswaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mahasiswa, err := m.mhsService.GetMahasiswaByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func (m *mahasiswaAPI) GetAllMahasiswa(c *gin.Context) {
	mahasiswa, err := m.mhsService.GetAllMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func (m *mahasiswaAPI) CreateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.mhsService.CreateMahasiswa(&mahasiswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Mahasiswa created successfully"})
}

func (m *mahasiswaAPI) UpdateMahasiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var mahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.mhsService.UpdateMahasiswa(id, &mahasiswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa updated successfully"})
}

func (m *mahasiswaAPI) DeleteMahasiswaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := m.mhsService.DeleteMahasiswa(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa deleted successfully"})
}
