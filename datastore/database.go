package datastore

import (
	"github.com/HashirMuhammad/todo-app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	err := d.conn.Where("id = ? ", Id).Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, err
}

func (d Database) Update(ID int) error {
	err := d.conn.Model(models.TodoItem{Id: ID}).Update(map[string]interface{}{"completed": true}).Error
	return err
}

func (d Database) Delete(item models.TodoItem, ID int) error {
	err := d.conn.Delete(&item, ID).Error
	return err
}

func (d Database) GetTodoItems(completed bool) ([]models.TodoItem, error) {
	var todo []models.TodoItem
	if err := d.conn.Where("completed = ?", completed).Find(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}
