package Antrian

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewPasien(c *gin.Context) {
	var pasien []Models.Pasien
	err := Models.PasienGetAll(&pasien)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":   "List App",
			"pasiens": pasien,
		}

		c.HTML(http.StatusOK, "index.html", data)
	}
}

func GetPasiens(c *gin.Context) {
	var pasien []Models.Pasien
	err := Models.PasienGetAll(&pasien)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pasien)
	}
}

func CreatePasien(c *gin.Context) {
	var pasien Models.Pasien
	c.BindJSON(&pasien)

	err := Models.PasienCreate(&pasien)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pasien)
	}
}

func ReadPasien(c *gin.Context) {
	var pasien Models.Pasien

	id := c.Params.ByName("id")
	err := Models.PasienFindOne(&pasien, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pasien)
	}
}

func UpdatePasien(c *gin.Context) {
	var pasien Models.Pasien

	id := c.Params.ByName("id")
	err := Models.PasienFindOne(&pasien, id)
	if err != nil {
		c.JSON(http.StatusNotFound, pasien)
	}

	c.BindJSON(&pasien)
	err = Models.PasienUpdateOne(&pasien, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pasien)
	}
}

func DeletePasien(c *gin.Context) {
	var pasien Models.Pasien

	id := c.Params.ByName("id")
	err := Models.PasienDeleteOne(&pasien, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
