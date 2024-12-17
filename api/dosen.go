package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DosenAPI interface {
	AddDosen(c *gin.Context)
	GetDosenByID(c *gin.Context)
	GetAllDosen(c *gin.Context)
}

type dosenAPI struct {
	dosenService services.DosenService
}

func NewDosenAPI(dosenRepo services.DosenService) *dosenAPI {
	return &dosenAPI{dosenRepo}
}

func (d *dosenAPI) AddDosen(c *gin.Context) {
	var dosen models.Dosen
	if err := c.ShouldBindJSON(&dosen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := d.dosenService.AddDosen(dosen); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Dosen created successfully"})
}

func (d *dosenAPI) GetDosenByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	dosen, err := d.dosenService.GetDosenByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dosen)
}

func (d *dosenAPI) GetAllDosen(c *gin.Context) {
	dosen, err := d.dosenService.GetAllDosen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dosen)
}
