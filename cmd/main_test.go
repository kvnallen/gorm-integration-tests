package main

import (
	"gorm-integration/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTodos(t *testing.T) {

	t.Run("when /POST users, should create a new one", func(t *testing.T) {
		// arrange
		db := createDb()
		handler := RegisterRoutes(db)

		// act
		res, err := http.NewRequest(http.MethodGet, "/create-todo", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, res)

		var todos []models.Todo
		db.Find(&todos)

		// assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 1, len(todos))
	})
}

func createDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Todo{})
	return db
}
