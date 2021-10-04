package datastore_test

import (
	"github.com/HashirMuhammad/todo-app/datastore"
	"github.com/HashirMuhammad/todo-app/models"
	"testing"
)

func TestDatabase_NewDatabase(t *testing.T) {
	_, err := datastore.NewDatabase()
	if err != nil {
		t.Error("error creating database")
	}
}

func TestDatabase_CreateItem(t *testing.T) {
	db, err := datastore.NewDatabase()
	if err != nil {
		t.Error("error creating database")
	}
	item := models.TodoItem{
		Id:          202,
		Description: "for testing",
		Completed:   false,
	}
	err = db.CreateItem(item)
	if err != nil {
		t.Error(err)
	}
}

func TestDatabase_GetItemByID(t *testing.T) {
	db, err := datastore.NewDatabase()
	if err != nil {
		t.Error("error creating database")
	}

	item, err := db.GetItemByID(202)
	if err != nil {
		t.Error(err)
	}

	if item.Id != 202 {
		t.Error("id not matched")
	}
}

func TestDatabase_Update(t *testing.T) {
	db, err := datastore.NewDatabase()
	if err != nil {
		t.Error("error creating database")
	}
	item, err := db.GetItemByID(202)
	if err != nil {
		t.Error(err)
	}

	item.Id = 202
	item.Description = "Its been updated"
	item.Completed = true
}

func TestDatabase_Delete(t *testing.T) {
	db, err := datastore.NewDatabase()
	if err != nil {
		t.Error("error creating database")
	}
	//item, err := db.GetItemByID(202)
	//if err != nil {
	//	t.Error(err)
	//}
	err = db.Delete(models.TodoItem{}, 202)
	if err != nil {
		t.Error(err)
	}

	//if (item.Id == 202) {
	//	t.Error("id not deleted")
	//}
}

func TestDatabase_GetTodoItems(t *testing.T) {
	db, err := datastore.NewDatabase()
	if err != nil {
		t.Error("error creating database")
	}
	_, err = db.GetTodoItems(true)
	if err != nil {
		t.Error(err)
	}

}
