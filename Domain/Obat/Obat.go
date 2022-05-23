package Obat

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewObat(c *gin.Context) {
	var obat []Models.Obat
	err := Models.ObatGetAll(&obat)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":  "List App",
			"obats": obat,
		}

		c.HTML(http.StatusOK, "obat.html", data)
	}
}

func GetObats(c *gin.Context) {
	var obat []Models.Obat
	err := Models.ObatGetAll(&obat)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, obat)
	}
}

func CreateObat(c *gin.Context) {
	var obat Models.Obat
	c.BindJSON(&obat)

	err := Models.ObatCreate(&obat)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, obat)
	}
}

func ReadObat(c *gin.Context) {
	var obat Models.Obat

	id := c.Params.ByName("id")
	err := Models.ObatFindOne(&obat, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, obat)
	}
}

func UpdateObat(c *gin.Context) {
	var obat Models.Obat

	id := c.Params.ByName("id")
	err := Models.ObatFindOne(&obat, id)
	if err != nil {
		c.JSON(http.StatusNotFound, obat)
	}

	c.BindJSON(&obat)
	err = Models.ObatUpdateOne(&obat, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, obat)
	}
}

func DeleteObat(c *gin.Context) {
	var obat Models.Obat

	id := c.Params.ByName("id")
	err := Models.ObatDeleteOne(&obat, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
