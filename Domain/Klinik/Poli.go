package Klinik

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewPoli(c *gin.Context) {
	var poli []Models.Poli
	var kliniks []Models.Klinik

	err := Models.PoliGetAll(&poli)
	Models.KlinikGetAll(&kliniks)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"polis":   poli,
			"kliniks": kliniks,
		}

		c.HTML(http.StatusOK, "poli.html", data)
	}
}

func GetPolis(c *gin.Context) {
	var poli []Models.Poli
	err := Models.PoliGetAll(&poli)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, poli)
	}
}

func CreatePoli(c *gin.Context) {
	var poli Models.Poli
	c.BindJSON(&poli)

	err := Models.PoliCreate(&poli)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, poli)
	}
}

func ReadPoli(c *gin.Context) {
	var poli Models.Poli

	id := c.Params.ByName("id")
	err := Models.PoliFindOne(&poli, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, poli)
	}
}

func UpdatePoli(c *gin.Context) {
	var poli Models.Poli

	id := c.Params.ByName("id")
	err := Models.PoliFindOne(&poli, id)
	if err != nil {
		c.JSON(http.StatusNotFound, poli)
	}

	c.BindJSON(&poli)
	err = Models.PoliUpdateOne(&poli, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, poli)
	}
}

func DeletePoli(c *gin.Context) {
	var poli Models.Poli

	id := c.Params.ByName("id")
	err := Models.PoliDeleteOne(&poli, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
