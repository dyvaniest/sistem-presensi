package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PresensiAPI interface {
	RecordPresensi(c *gin.Context)
	InsertPresensi(c *gin.Context)
	GetPresensiByID(c *gin.Context)
	GetAllPresensi(c *gin.Context)
	UpdatePresensi(c *gin.Context)
	DeletePresensiByID(c *gin.Context)
}

type presensiAPI struct {
	presensiService services.PresensiService
}

// new presensi api
func NewPresensiAPI(presensiRepo services.PresensiService) *presensiAPI {
	return &presensiAPI{presensiRepo}
}

func (p *presensiAPI) RecordPresensi(c *gin.Context) {
	uid := c.Query("uid")
	pertemuanID, err := strconv.Atoi(c.Query("pertemuan_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pertemuan_id"})
		return
	}

	message, err := p.presensiService.RecordPresensi(uid, pertemuanID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (p *presensiAPI) InsertPresensi(c *gin.Context) {
	var presensi models.Presensi
	if err := c.ShouldBindJSON(&presensi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.presensiService.InsertPresensi(&presensi); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Presensi created successfully"})
}

func (p *presensiAPI) GetPresensiByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	presensi, err := p.presensiService.GetPresensiByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, presensi)
}

func (p *presensiAPI) GetAllPresensi(c *gin.Context) {
	presensi, err := p.presensiService.GetAllPresensi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, presensi)
}

func (p *presensiAPI) UpdatePresensi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var presensi models.Presensi
	if err := c.ShouldBindJSON(&presensi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.presensiService.UpdatePresensi(id, presensi); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Presensi updated successfully"})
}

func (p *presensiAPI) DeletePresensiByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := p.presensiService.DeletePresensiByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Presensi deleted successfully"})
}
