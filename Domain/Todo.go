package Domain

import (
	"golang-klinik/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewTodo(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAll(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title": "TodoList App",
			"todos": todo,
		}

		c.HTML(http.StatusOK, "todo.html", data)
	}
}

func GetTodos(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAll(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)

	err := Models.Create(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func ReadTodo(c *gin.Context) {
	var todo Models.Todo

	id := c.Params.ByName("id")
	err := Models.FindOne(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo Models.Todo

	id := c.Params.ByName("id")
	err := Models.FindOne(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}

	c.BindJSON(&todo)
	err = Models.UpdateOne(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo Models.Todo

	id := c.Params.ByName("id")
	err := Models.DeleteOne(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id:" + id: "deleted"})
	}
}
