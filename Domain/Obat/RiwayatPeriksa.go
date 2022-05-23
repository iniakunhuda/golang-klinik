package Obat

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewRiwayatPeriksa(c *gin.Context) {
	var riwayat_periksa []Models.RiwayatPeriksa
	err := Models.RiwayatPeriksaGetAll(&riwayat_periksa)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":            "List App",
			"riwayat_periksas": riwayat_periksa,
		}

		c.HTML(http.StatusOK, "index.html", data)
	}
}

func GetRiwayatPeriksas(c *gin.Context) {
	var riwayat_periksa []Models.RiwayatPeriksa
	err := Models.RiwayatPeriksaGetAll(&riwayat_periksa)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, riwayat_periksa)
	}
}

func CreateRiwayatPeriksa(c *gin.Context) {
	var riwayat_periksa Models.RiwayatPeriksa
	c.BindJSON(&riwayat_periksa)

	err := Models.RiwayatPeriksaCreate(&riwayat_periksa)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, riwayat_periksa)
	}
}

func ReadRiwayatPeriksa(c *gin.Context) {
	var riwayat_periksa Models.RiwayatPeriksa

	id := c.Params.ByName("id")
	err := Models.RiwayatPeriksaFindOne(&riwayat_periksa, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, riwayat_periksa)
	}
}

func UpdateRiwayatPeriksa(c *gin.Context) {
	var riwayat_periksa Models.RiwayatPeriksa

	id := c.Params.ByName("id")
	err := Models.RiwayatPeriksaFindOne(&riwayat_periksa, id)
	if err != nil {
		c.JSON(http.StatusNotFound, riwayat_periksa)
	}

	c.BindJSON(&riwayat_periksa)
	err = Models.RiwayatPeriksaUpdateOne(&riwayat_periksa, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, riwayat_periksa)
	}
}

func DeleteRiwayatPeriksa(c *gin.Context) {
	var riwayat_periksa Models.RiwayatPeriksa

	id := c.Params.ByName("id")
	err := Models.RiwayatPeriksaDeleteOne(&riwayat_periksa, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
