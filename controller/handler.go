package controller

import (
	"encoding/json"
	"fmt"
	"github.com/HashirMuhammad/todo-app/datastore"
	"github.com/HashirMuhammad/todo-app/models"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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

	// Test if the TodoItem exist in DB
	todo, err := c.Db.GetItemByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"updated": false, "error": "Record Not Found"}`)
		return
	}

	completed, err := strconv.ParseBool(r.FormValue("completed"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"updated": false, "error": "Bad Request"}`)
		return
	}

	todo.Completed = completed
	err = c.Db.Update(todo, id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"updated": false, "error": "error"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"updated": true}`)
}

func (c Controller) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Test if the TodoItem exist in DB
	todo, err := c.Db.GetItemByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
		return
	}
		err = c.Db.Delete(todo, id)
	if  err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"updated": false, "error": "error"}`)
		return
	}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
}
