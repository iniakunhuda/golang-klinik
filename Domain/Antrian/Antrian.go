package Antrian

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAntrian(c *gin.Context) {
	var antrian []Models.Antrian
	err := Models.AntrianGetAll(&antrian)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":    "List App",
			"antrians": antrian,
		}

		c.HTML(http.StatusOK, "index.html", data)
	}
}

func GetAntrians(c *gin.Context) {
	var antrian []Models.Antrian
	err := Models.AntrianGetAll(&antrian)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, antrian)
	}
}

func CreateAntrian(c *gin.Context) {
	var antrian Models.Antrian
	c.BindJSON(&antrian)

	err := Models.AntrianCreate(&antrian)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, antrian)
	}
}

func ReadAntrian(c *gin.Context) {
	var antrian Models.Antrian

	id := c.Params.ByName("id")
	err := Models.AntrianFindOne(&antrian, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, antrian)
	}
}

func UpdateAntrian(c *gin.Context) {
	var antrian Models.Antrian

	id := c.Params.ByName("id")
	err := Models.AntrianFindOne(&antrian, id)
	if err != nil {
		c.JSON(http.StatusNotFound, antrian)
	}

	c.BindJSON(&antrian)
	err = Models.AntrianUpdateOne(&antrian, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, antrian)
	}
}

func DeleteAntrian(c *gin.Context) {
	var antrian Models.Antrian

	id := c.Params.ByName("id")
	err := Models.AntrianDeleteOne(&antrian, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
