package main

import (
	"encoding/json"
	"gorm-integration/models"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type app struct {
	gorm.DB
}

func main() {
	dsn := "dev:dev@tcp(localhost:3306)/dev-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", RegisterRoutes(db))
}

func RegisterRoutes(db *gorm.DB) http.Handler {
	app := &app{
		DB: *db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", app.TodoListHandler)
	mux.HandleFunc("/create-todo", app.CreateTodoHandler)
	return mux
}

func (a *app) TodoListHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Searching for all todos")
	var todos []models.Todo
	a.DB.Find(&todos)
	todosJSON, _ := json.Marshal(&todos)
	w.Header().Add("Content-Type", "application/json")
	w.Write(todosJSON)
}

func (a *app) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := a.DB
	db.AutoMigrate(&models.Todo{})
	db.Create(&models.Todo{
		Title:       "Use GORM",
		Description: "create integration tests",
		Roles:       []byte(`[{ "name": "admin", "active": true }]`),
	})
}
