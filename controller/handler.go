package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/HashirMuhammad/todo-app/datastore"
	"github.com/HashirMuhammad/todo-app/models"
	"github.com/gorilla/mux"
)

type Controller struct {
	Db datastore.Database
}

func (c Controller) CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	todo := models.TodoItem{Description: description, Completed: false}
	err := c.Db.CreateItem(todo)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err.Error()))
	}
	json.NewEncoder(w).Encode("item craeted sucessfully ")
}

func (c Controller) UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := c.Db.Update(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"updated": false, "error": "error"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"updated": true}`)
}

func (c Controller) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err.Error()))
	}
	// Test if the TodoItem exist in DB
	w.Header().Set("Content-Type", "application/json")

	todo, err := c.Db.GetItemByID(id)
	if err != nil {
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
		return
	}

	if err = c.Db.Delete(todo, id); err != nil {
		json.NewEncoder(w).Encode(`{"updated": false, "error": "error"}`)
		return
	}

	json.NewEncoder(w).Encode(`{"deleted": true}`)
}

func (c Controller) GetAllItems(w http.ResponseWriter, r *http.Request) {
	form, err := strconv.ParseBool(r.FormValue("completed"))

	items, err := c.Db.GetTodoItems(form)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
