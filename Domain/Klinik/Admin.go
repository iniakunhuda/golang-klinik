package Klinik

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAdmin(c *gin.Context) {
	var admin []Models.Admin
	err := Models.AdminGetAll(&admin)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":  "List App",
			"admins": admin,
		}

		c.HTML(http.StatusOK, "index.html", data)
	}
}

func GetAdmins(c *gin.Context) {
	var admin []Models.Admin
	err := Models.AdminGetAll(&admin)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

func CreateAdmin(c *gin.Context) {
	var admin Models.Admin
	c.BindJSON(&admin)

	err := Models.AdminCreate(&admin)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

func ReadAdmin(c *gin.Context) {
	var admin Models.Admin

	id := c.Params.ByName("id")
	err := Models.AdminFindOne(&admin, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

func UpdateAdmin(c *gin.Context) {
	var admin Models.Admin

	id := c.Params.ByName("id")
	err := Models.AdminFindOne(&admin, id)
	if err != nil {
		c.JSON(http.StatusNotFound, admin)
	}

	c.BindJSON(&admin)
	err = Models.AdminUpdateOne(&admin, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

func DeleteAdmin(c *gin.Context) {
	var admin Models.Admin

	id := c.Params.ByName("id")
	err := Models.AdminDeleteOne(&admin, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
