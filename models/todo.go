package models

import (
	"time"
	"todo/dao"
)

type Todo struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string    `json:"title" gorm:"not null"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;not null"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime;not null"`
	DeleteTime time.Time `json:"delete_time"`
}

func CreateTodo(todo *Todo) error {
	err := dao.Db.Create(todo).Error
	return err
}

func DeleteTodoById(id string) error {
	var todo *Todo
	if err := dao.Db.Where("id=?", id).First(&todo).Error; err != nil {
		panic(err)
		return err
	}
	if err := dao.Db.Model(&Todo{}).Where("id=?", id).Update("delete_time", time.Now()).Error; err != nil {
		panic(err)
		return err
	}
	return nil
}

func UpdateTodoById(id uint) {

}

func GetAllTodo() {

}
