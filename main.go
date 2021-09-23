package main

import (
	"io"
	"net/http"

	"github.com/HashirMuhammad/todo-app/controller"
	"github.com/HashirMuhammad/todo-app/datastore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main() {
	db, err := datastore.NewDatabase()
	if err != nil {
		panic(err)
	}

	ctrl := controller.Controller{Db: db}

	// defer db.Close()

	// db.Debug().DropTableIfExists(&TodoItemModel{})
	// db.Debug().AutoMigrate(&TodoItemModel{})

	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/items", ctrl.GetAllItems).Methods("GET")
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
