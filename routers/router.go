package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// gin框架 加载静态文件
	r.Static("/static", "static")

	// gin框架 解析模板
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("/v1")
	{
		/*-------------------待办事项--------------*/
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 查看某一个待办事项 (URL参数 id)
		v1Group.GET("/todo/:id", controller.GetTodoById)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
