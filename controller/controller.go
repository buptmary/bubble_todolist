package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	url --> controller --> logic --> model
	请求 -->  控制器     --> 业务逻辑 --> 模型的CRUD
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 发送请求到这里
	// 1. 从请求中获取数据
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "get todo fail!"})
		return
	}

	// 2. 存入数据库
	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 3. 返回相应
	c.JSON(http.StatusOK, todo)
}

func GetTodoById(c *gin.Context) {
	// 查询todo表中指定id数据
	id := c.Param("id")
	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func GetTodoList(c *gin.Context) {
	// 查询todo表中所有数据
	todoList, err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 返回todoList数据
	c.JSON(http.StatusOK, todoList)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 将数据绑定JSON
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	// 获取要删除的id
	id := c.Param("id")

	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "delete success!"})
}
