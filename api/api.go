package api

import (
	repo "sistem-presensi/repository"
	"sistem-presensi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPI       UserAPI
	PresensiAPI   PresensiAPI
	MahasiswaAPI  MahasiswaAPI
	DosenAPI      DosenAPI
	JadwalAPI     JadwalAPI
	MataKuliahAPI MataKuliahAPI
	PertemuanAPI  PertemuanAPI
}

// RunServer initializes all routes and returns a Gin engine
func RunServer(router *gin.Engine, db *gorm.DB) *gin.Engine {
	// Repository Initialization
	userRepo := repo.NewUserRepo(db)
	sessionRepo := repo.NewSessionsRepo(db)
	presensiRepo := repo.NewPresensiRepository(db)
	mahasiswaRepo := repo.NewMahasiswaRepository(db)
	dosenRepo := repo.NewDosenRepo(db)
	jadwalRepo := repo.NewJadwalRespository(db)
	mataKuliahRepo := repo.NewMataKuliahRepository(db)
	pertemuanRepo := repo.NewPertemuanRepo(db)

	// Service Initialization
	userService := services.NewUserService(userRepo, sessionRepo)
	presensiService := services.NewPresensiService(presensiRepo)
	mahasiswaService := services.NewMahasiswaService(mahasiswaRepo)
	dosenService := services.NewDosenService(dosenRepo)
	jadwalService := services.NewJadwalService(jadwalRepo)
	mataKuliahService := services.NewMataKuliahService(mataKuliahRepo)
	pertemuanService := services.NewPertemuanService(pertemuanRepo)

	// API Handlers Initialization
	userAPIHandler := NewUserAPI(userService)
	presensiAPIHandler := NewPresensiAPI(presensiService)
	mahasiswaAPIHandler := NewMahasiswaAPI(mahasiswaService)
	dosenAPIHandler := NewDosenAPI(dosenService)
	jadwalAPIHandler := NewJadwalAPI(jadwalService)
	mataKuliahAPIHandler := NewMataKuliahAPI(mataKuliahService)
	pertemuanAPIHandler := NewPertemuanAPI(pertemuanService)

	apiHandler := APIHandler{
		UserAPI:       userAPIHandler,
		PresensiAPI:   presensiAPIHandler,
		MahasiswaAPI:  mahasiswaAPIHandler,
		DosenAPI:      dosenAPIHandler,
		JadwalAPI:     jadwalAPIHandler,
		MataKuliahAPI: mataKuliahAPIHandler,
		PertemuanAPI:  pertemuanAPIHandler,
	}

	// API Versioning
	version := router.Group("/api/v1")
	{
		user := version.Group("/user")
		{
			user.POST("/register", apiHandler.UserAPI.Register)
			user.POST("/login", apiHandler.UserAPI.Login)
		}

		// Presensi routes
		presensi := version.Group("/presensi")
		{
			presensi.POST("/record-presensi", apiHandler.PresensiAPI.RecordPresensi)
			presensi.GET("/", apiHandler.PresensiAPI.GetAllPresensi)
			presensi.POST("/", apiHandler.PresensiAPI.InsertPresensi)
			presensi.GET("/:id", apiHandler.PresensiAPI.GetPresensiByID)
			presensi.PUT("/:id", apiHandler.PresensiAPI.UpdatePresensi)
			presensi.DELETE("/:id", apiHandler.PresensiAPI.DeletePresensiByID)
			presensi.GET("/record-presensi", apiHandler.PresensiAPI.RecordPresensi)
		}

		// Mahasiswa routes
		mahasiswa := version.Group("/mahasiswa")
		{
			mahasiswa.GET("/", apiHandler.MahasiswaAPI.GetAllMahasiswa)
			mahasiswa.POST("/", apiHandler.MahasiswaAPI.CreateMahasiswa)
			mahasiswa.GET("/:id", apiHandler.MahasiswaAPI.GetMahasiswaByID)
			mahasiswa.PUT("/:id", apiHandler.MahasiswaAPI.UpdateMahasiswa)
			mahasiswa.DELETE("/:id", apiHandler.MahasiswaAPI.DeleteMahasiswaByID)
		}

		// Dosen routes
		dosen := version.Group("/dosen")
		{
			dosen.GET("/", apiHandler.DosenAPI.GetAllDosen)
			dosen.POST("/", apiHandler.DosenAPI.AddDosen)
			dosen.GET("/:id", apiHandler.DosenAPI.GetDosenByID)
		}

		// Jadwal routes
		jadwal := version.Group("/jadwal")
		{
			jadwal.POST("/", apiHandler.JadwalAPI.AddJadwal)
			jadwal.GET("/:id", apiHandler.JadwalAPI.GetJadwalByID)
			jadwal.PUT("/:id", apiHandler.JadwalAPI.UpdateJadwal)
			jadwal.DELETE("/:id", apiHandler.JadwalAPI.DeleteJadwalByID)
		}

		// Mata Kuliah routes
		mataKuliah := version.Group("/mata-kuliah")
		{
			mataKuliah.GET("/:id", apiHandler.MataKuliahAPI.GetMataKuliahByID)
			mataKuliah.PUT("/:id", apiHandler.MataKuliahAPI.UpdateMataKuliah)
			mataKuliah.DELETE("/:id", apiHandler.MataKuliahAPI.DeleteMataKuliahByID)
		}

		pertemuan := version.Group("/pertemuan")
		{
			pertemuan.GET("/", apiHandler.PertemuanAPI.GetAllPertemuan)
			pertemuan.GET("/:id", apiHandler.PertemuanAPI.GetPertemuanByID)
			pertemuan.POST("/", apiHandler.PertemuanAPI.AddPertemuan)
			pertemuan.PUT("/:id", apiHandler.PertemuanAPI.UpdatePertemuan)
			pertemuan.DELETE("/:id", apiHandler.PertemuanAPI.DeletePertemuanByID)
		}
	}

	return router
}
