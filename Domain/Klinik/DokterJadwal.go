package Klinik

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewDokterJadwal(c *gin.Context) {
	var dokter_jadwal []Models.DokterJadwal
	var dokter []Models.Dokter
	var klinik []Models.Klinik
	var poli []Models.Poli

	err := Models.DokterJadwalGetAll(&dokter_jadwal)

	Models.DokterGetAll(&dokter)
	Models.KlinikGetAll(&klinik)
	Models.PoliGetAll(&poli)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":          "List App",
			"dokter_jadwals": dokter_jadwal,
			"dokters":        dokter,
			"kliniks":        klinik,
			"polis":          poli,
		}

		c.HTML(http.StatusOK, "jadwal_dokter.html", data)
	}
}

func GetDokterJadwals(c *gin.Context) {
	var dokter_jadwal []Models.DokterJadwal
	err := Models.DokterJadwalGetAll(&dokter_jadwal)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter_jadwal)
	}
}

func CreateDokterJadwal(c *gin.Context) {
	var dokter_jadwal Models.DokterJadwal
	c.BindJSON(&dokter_jadwal)

	err := Models.DokterJadwalCreate(&dokter_jadwal)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter_jadwal)
	}
}

func ReadDokterJadwal(c *gin.Context) {
	var dokter_jadwal Models.DokterJadwal

	id := c.Params.ByName("id")
	err := Models.DokterJadwalFindOne(&dokter_jadwal, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter_jadwal)
	}
}

func UpdateDokterJadwal(c *gin.Context) {
	var dokter_jadwal Models.DokterJadwal

	id := c.Params.ByName("id")
	err := Models.DokterJadwalFindOne(&dokter_jadwal, id)
	if err != nil {
		c.JSON(http.StatusNotFound, dokter_jadwal)
	}

	c.BindJSON(&dokter_jadwal)
	err = Models.DokterJadwalUpdateOne(&dokter_jadwal, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter_jadwal)
	}
}

func DeleteDokterJadwal(c *gin.Context) {
	var dokter_jadwal Models.DokterJadwal

	id := c.Params.ByName("id")
	err := Models.DokterJadwalDeleteOne(&dokter_jadwal, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
