package Klinik

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewDokter(c *gin.Context) {
	var dokter []Models.Dokter
	err := Models.DokterGetAll(&dokter)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":   "List App",
			"dokters": dokter,
		}

		c.HTML(http.StatusOK, "dokter.html", data)
	}
}

func GetDokters(c *gin.Context) {
	var dokter []Models.Dokter
	err := Models.DokterGetAll(&dokter)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter)
	}
}

func CreateDokter(c *gin.Context) {
	var dokter Models.Dokter
	c.BindJSON(&dokter)

	err := Models.DokterCreate(&dokter)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter)
	}
}

func ReadDokter(c *gin.Context) {
	var dokter Models.Dokter

	id := c.Params.ByName("id")
	err := Models.DokterFindOne(&dokter, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter)
	}
}

func UpdateDokter(c *gin.Context) {
	var dokter Models.Dokter

	id := c.Params.ByName("id")
	err := Models.DokterFindOne(&dokter, id)
	if err != nil {
		c.JSON(http.StatusNotFound, dokter)
	}

	c.BindJSON(&dokter)
	err = Models.DokterUpdateOne(&dokter, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dokter)
	}
}

func DeleteDokter(c *gin.Context) {
	var dokter Models.Dokter

	id := c.Params.ByName("id")
	err := Models.DokterDeleteOne(&dokter, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
