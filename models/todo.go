package models

import (
	"bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo 增删改查操作
*/

func InitModel() {
	err := dao.DB.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}
}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoById(id string) (todo *Todo, err error) {
	// 对Todo指针初始化
	todo = new(Todo)

	if err := dao.DB.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id =?", id).Delete(&Todo{}).Error
	return
}
