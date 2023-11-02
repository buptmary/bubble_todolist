package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

// lesson27 企业级项目结构拆分

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;

	// 连接数据库
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}

	// 模型绑定
	models.InitModel()

	// 注册路由
	r := routers.SetupRouter()
	r.Run()
}
