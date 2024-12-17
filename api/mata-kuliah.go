package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MataKuliahAPI interface {
	GetMataKuliahByID(c *gin.Context)
	UpdateMataKuliah(c *gin.Context)
	DeleteMataKuliahByID(c *gin.Context)
}

type mataKuliahAPI struct {
	matkulService services.MataKuliahService
}

func NewMataKuliahAPI(matkulRepo services.MataKuliahService) *mataKuliahAPI {
	return &mataKuliahAPI{matkulRepo}
}

func (mk *mataKuliahAPI) GetMataKuliahByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	matkul, err := mk.matkulService.GetMataKuliahByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, matkul)
}

func (mk *mataKuliahAPI) UpdateMataKuliah(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var mataKuliah models.MataKuliah
	if err := c.ShouldBindJSON(&mataKuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mk.matkulService.UpdateMataKuliah(id, mataKuliah); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mata Kuliah updated successfully"})
}

func (mk *mataKuliahAPI) DeleteMataKuliahByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := mk.matkulService.DeleteMataKuliahByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mata Kuliah deleted successfully"})
}
