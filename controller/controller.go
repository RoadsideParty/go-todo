package controller

import (
	"github.com/gin-gonic/gin"
	"todo/models"
	"todo/response"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		response.Fail(c, "添加失败，缺少参数")
		return
	}
	if todo.Title == "" {
		response.Fail(c, "title为必填项")
		return
	}
	if err := models.CreateTodo(&todo); err != nil {
		response.Fail(c, "数据库错误")
		return
	}
	response.Ok(c, &todo)
}

func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteTodoById(id)
	if err != nil {
		response.Fail(c, "数据库错误")
		return
	}
	response.Ok(c, nil)
}
