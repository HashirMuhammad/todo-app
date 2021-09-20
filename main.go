package main

import (
	 "github.com/HashirMuhammad/todo-app/controller"
	"github.com/HashirMuhammad/todo-app/datastore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
	"io"
	"net/http"
)



//

//
//func (n dat) GetCompletedItems(w http.ResponseWriter, r *http.Request) {
//	completedTodoItems := n.GetTodoItems(true)
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(completedTodoItems)
//}
//
//func (n dat) GetIncompleteItems(w http.ResponseWriter, r *http.Request) {
//
//	IncompleteTodoItems := n.GetTodoItems(false)
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(IncompleteTodoItems)
//}
//
//func (n dat) GetTodoItems(completed bool) []TodoItemModel {
//	var todos []TodoItemModel
//	err := n.db.Where("completed = ?", completed).Find(&todos).Error
//	if err != nil {
//		return nil
//	}
//	return todos
//}

func main() {
	db, err := datastore.NewDatabase()
	if err != nil {
		panic(err)
	}

	ctrl := controller.Controller{Db: db}

	//defer db.Close()

	//db.Debug().DropTableIfExists(&TodoItemModel{})
	//db.Debug().AutoMigrate(&TodoItemModel{})

	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	//router.HandleFunc("/todo-completed", n.GetCompletedItems).Methods("GET")
	//router.HandleFunc("/todo-incomplete", n.GetIncompleteItems).Methods("GET")
	router.HandleFunc("/todo", ctrl.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", ctrl.UpdateItem).Methods("PATCH")
	router.HandleFunc("/todo/{id}", ctrl.DeleteItem).Methods("DELETE")

	err = http.ListenAndServe(":8000", router)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8000", handler)

	if err != nil {
		return
	}

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		return
	}
}
