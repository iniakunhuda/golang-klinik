package Klinik

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewKlinik(c *gin.Context) {
	var klinik []Models.Klinik
	err := Models.KlinikGetAll(&klinik)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":   "List App",
			"kliniks": klinik,
		}

		c.HTML(http.StatusOK, "klinik.html", data)
	}
}

func GetKliniks(c *gin.Context) {
	var klinik []Models.Klinik
	err := Models.KlinikGetAll(&klinik)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, klinik)
	}
}

func CreateKlinik(c *gin.Context) {
	var klinik Models.Klinik
	c.BindJSON(&klinik)

	err := Models.KlinikCreate(&klinik)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, klinik)
	}
}

func ReadKlinik(c *gin.Context) {
	var klinik Models.Klinik

	id := c.Params.ByName("id")
	err := Models.KlinikFindOne(&klinik, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, klinik)
	}
}

func UpdateKlinik(c *gin.Context) {
	var klinik Models.Klinik

	id := c.Params.ByName("id")
	err := Models.KlinikFindOne(&klinik, id)
	if err != nil {
		c.JSON(http.StatusNotFound, klinik)
	}

	c.BindJSON(&klinik)
	err = Models.KlinikUpdateOne(&klinik, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, klinik)
	}
}

func DeleteKlinik(c *gin.Context) {
	var klinik Models.Klinik

	id := c.Params.ByName("id")
	err := Models.KlinikDeleteOne(&klinik, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
