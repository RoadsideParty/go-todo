package controller

import (
	"github.com/gin-gonic/gin"
	"todo/models"
	"todo/response"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if todo.Title == "" {
		response.Fail(c, "title为必填项")
		return
	}
	if err := models.CreateTodo(&todo); err != nil {
		response.Fail(c, "添加失败")
		return
	}
	response.Ok(c, &todo)
}
