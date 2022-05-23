package Routes

import (
	"golang-klinik/Domain"
	"golang-klinik/Domain/Antrian"
	"golang-klinik/Domain/Klinik"
	"golang-klinik/Domain/Obat"
	"path/filepath"

	"github.com/Masterminds/sprig"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// HOME
	layouts, err := filepath.Glob(templatesDir + "/layouts/default.html")
	if err != nil {
		panic(err.Error())
	}
	users, err := filepath.Glob(templatesDir + "/includes/home/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, user := range users {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, user)
		r.AddFromFilesFuncs(filepath.Base(user), sprig.FuncMap(), files...)
	}

	// Admin

	layouts_admin, err := filepath.Glob(templatesDir + "/layouts/admin.html")
	if err != nil {
		panic(err.Error())
	}
	admins, err := filepath.Glob(templatesDir + "/includes/admin/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, admin := range admins {
		layoutCopy := make([]string, len(layouts_admin))
		copy(layoutCopy, layouts_admin)
		files := append(layoutCopy, admin)
		r.AddFromFilesFuncs(filepath.Base(admin), sprig.FuncMap(), files...)
	}

	return r
}

func SetRoutes() *gin.Engine {
	r := gin.Default()

	//-- setup view
	r.HTMLRender = loadTemplates("./Views")
	r.GET("/", Domain.ViewIndex)
	r.GET("account/login", Domain.ViewAccountLogin)
	r.POST("account/login", Domain.AccountLogin)
	r.GET("account/register", Domain.ViewAccountRegister)
	r.POST("account/register", Domain.RegisterNewAccount)
	r.GET("account/logout", Domain.AccountLogout)

	r.POST("register", Domain.CreateRegister)
	r.GET("register/:klinik_id/:poli_id", Domain.ViewRegister)
	r.GET("todo", Domain.ViewTodo)

	admin := r.Group("admin")
	{

		admin.GET("/klinik", Klinik.ViewKlinik)
		admin.GET("/poli", Klinik.ViewPoli)
		admin.GET("/dokter", Klinik.ViewDokter)
		admin.GET("/jadwal_dokter", Klinik.ViewDokterJadwal)
		admin.GET("/obat", Obat.ViewObat)

	}

	/**
	====================================
	API SYSTEM V1
	====================================
	**/
	v1 := r.Group("v1")
	{
		todo := v1.Group("todo")
		{
			todo.GET("/", Domain.GetTodos)
			todo.POST("/", Domain.CreateTodo)
			todo.GET(":id", Domain.ReadTodo)
			todo.PUT(":id", Domain.UpdateTodo)
			todo.DELETE(":id", Domain.DeleteTodo)
		}

		/**
		DOMAIN KLINIK
		**/
		domain_klinik := v1.Group("klinik")
		{

			klinik := domain_klinik.Group("klinik")
			{
				klinik.GET("/", Klinik.GetKliniks)
				klinik.POST("/", Klinik.CreateKlinik)
				klinik.GET(":id", Klinik.ReadKlinik)
				klinik.PUT(":id", Klinik.UpdateKlinik)
				klinik.DELETE(":id", Klinik.DeleteKlinik)
			}

			admin := domain_klinik.Group("admin")
			{
				admin.GET("/", Klinik.GetAdmins)
				admin.POST("/", Klinik.CreateAdmin)
				admin.GET(":id", Klinik.ReadAdmin)
				admin.PUT(":id", Klinik.UpdateAdmin)
				admin.DELETE(":id", Klinik.DeleteAdmin)
			}

			dokter := domain_klinik.Group("dokter")
			{
				dokter.GET("/", Klinik.GetDokters)
				dokter.POST("/", Klinik.CreateDokter)
				dokter.GET(":id", Klinik.ReadDokter)
				dokter.PUT(":id", Klinik.UpdateDokter)
				dokter.DELETE(":id", Klinik.DeleteDokter)
			}

			poli := domain_klinik.Group("poli")
			{
				poli.GET("/", Klinik.GetPolis)
				poli.POST("/", Klinik.CreatePoli)
				poli.GET(":id", Klinik.ReadPoli)
				poli.PUT(":id", Klinik.UpdatePoli)
				poli.DELETE(":id", Klinik.DeletePoli)
			}

			dokter_jadwal := domain_klinik.Group("dokter_jadwal")
			{
				dokter_jadwal.GET("/", Klinik.GetDokterJadwals)
				dokter_jadwal.POST("/", Klinik.CreateDokterJadwal)
				dokter_jadwal.GET(":id", Klinik.ReadDokterJadwal)
				dokter_jadwal.PUT(":id", Klinik.UpdateDokterJadwal)
				dokter_jadwal.DELETE(":id", Klinik.DeleteDokterJadwal)
			}

		}

		/**
		DOMAIN ANTRI
		**/
		domain_antri := v1.Group("antrian")
		{

			pasien := domain_antri.Group("pasien")
			{
				pasien.GET("/", Antrian.GetPasiens)
				pasien.POST("/", Antrian.CreatePasien)
				pasien.GET(":id", Antrian.ReadPasien)
				pasien.PUT(":id", Antrian.UpdatePasien)
				pasien.DELETE(":id", Antrian.DeletePasien)
			}

			antrian := domain_antri.Group("antrian")
			{
				antrian.GET("/", Antrian.GetAntrians)
				antrian.POST("/", Antrian.CreateAntrian)
				antrian.GET(":id", Antrian.ReadAntrian)
				antrian.PUT(":id", Antrian.UpdateAntrian)
				antrian.DELETE(":id", Antrian.DeleteAntrian)
			}

		}

		/**
		DOMAIN OBAT
		**/
		domain_obat := v1.Group("obat")
		{

			obat := domain_obat.Group("obat")
			{
				obat.GET("/", Obat.GetObats)
				obat.POST("/", Obat.CreateObat)
				obat.GET(":id", Obat.ReadObat)
				obat.PUT(":id", Obat.UpdateObat)
				obat.DELETE(":id", Obat.DeleteObat)
			}

			riwayat_periksa := domain_obat.Group("riwayat_periksa")
			{
				riwayat_periksa.GET("/", Obat.GetRiwayatPeriksas)
				riwayat_periksa.POST("/", Obat.CreateRiwayatPeriksa)
				riwayat_periksa.GET(":id", Obat.ReadRiwayatPeriksa)
				riwayat_periksa.PUT(":id", Obat.UpdateRiwayatPeriksa)
				riwayat_periksa.DELETE(":id", Obat.DeleteRiwayatPeriksa)
			}

		}

	}

	return r
}
