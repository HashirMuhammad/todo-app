package main

import (
	"io"
	"net/http"
	"strconv"

	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func (n n) UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Test if the TodoItem exist in DB
	err := n.GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": false, "error": "Record Not Found"}`)
	} else {
		completed, _ := strconv.ParseBool(r.FormValue("completed"))
		todo := &TodoItemModel{}
		n.db.First(&todo, id)
		todo.Completed = completed
		n.db.Save(&todo)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": true}`)
	}
}

func (n n) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Test if the TodoItem exist in DB
	err := n.GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
	} else {
		todo := &TodoItemModel{}
		n.db.First(&todo, id)
		n.db.Delete(&todo)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
	}
}

func (n n) GetItemByID(Id int) bool {
	todo := &TodoItemModel{}
	result := n.db.First(&todo, Id)
	if result.Error != nil {
		return false
	}
	return true
}

func (n n) GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	completedTodoItems := n.GetTodoItems(true)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completedTodoItems)
}

func (n n) GetIncompleteItems(w http.ResponseWriter, r *http.Request) {

	IncompleteTodoItems := n.GetTodoItems(false)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IncompleteTodoItems)
}

func (n n) GetTodoItems(completed bool) []TodoItemModel {
	var todos []TodoItemModel
	err := n.db.Where("completed = ?", completed).Find(&todos).Error
	if err != nil {
		return nil
	}
	return todos
}

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

func (n n) CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	todo := &TodoItemModel{Description: description, Completed: false}
	n.db.Create(&todo)
	result := n.db.Last(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{})
	n := n{db: db}
	router := mux.NewRouter()
	router.HandleFunc("/healthz", n.Healthz).Methods("GET")
	router.HandleFunc("/todo-completed", n.GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", n.GetIncompleteItems).Methods("GET")
	router.HandleFunc("/todo", n.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", n.UpdateItem).Methods("PATCH")
	router.HandleFunc("/todo/{id}", n.DeleteItem).Methods("DELETE")

	err = http.ListenAndServe(":8000", router)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8000", handler)

	if err != nil {
		return
	}

}

type n struct {
	db *gorm.DB
}

func (n n) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		return
	}
}
