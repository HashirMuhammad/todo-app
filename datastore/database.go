package datastore

import (
	"github.com/HashirMuhammad/todo-app/models"
	"github.com/jinzhu/gorm"
)

type Database struct {
	conn *gorm.DB
}

func NewDatabase() (Database, error) {
	db, err := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")
	return Database{db}, err
}

func (d Database) CreateItem(todo models.TodoItem) error {
	err := d.conn.Create(&todo).Error
	return err
}

func (d Database) GetItemByID(Id int) (models.TodoItem, error) {
	todo := models.TodoItem{}
	err := d.conn.Select(&todo, Id).Exec("SELECT Id FROM TodoItem", &todo.Id).Error
	//if err != nil {
	//	return todo, err
	//}

	return todo, err
}

func (d Database) Update(item models.TodoItem, ID int) error {
	err := d.conn.Update(&item, ID).Error
	return err
}
func (d Database) Delete(item models.TodoItem, ID int) error  {
	err := d.conn.Delete(&item, ID).Error
	return err

}

