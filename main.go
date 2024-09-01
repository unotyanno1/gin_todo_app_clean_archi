package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"gin_todo_app_clean_archi/src/domain/models"
	"gin_todo_app_clean_archi/src/infra/database/repositories"
	"gin_todo_app_clean_archi/src/infra/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := database.ConnectionDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	todoRepo := repository.NewTodoRepository(db)

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r.Static("/static", "./static")
	r.LoadHTMLGlob("src/infra/http/public/*")

	r.GET("/index", func(c *gin.Context) {
		var todos []*models.Todo
		todos, err := todoRepo.List(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve todos"})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Todoアプリ　クリーンアーキテクチャー版",
			"todos": todos,
		})
	})

	r.GET("/todos/edit", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not invalid parameter"})
			return
		}
		var todo *models.Todo
		todo, _ = todoRepo.GetByID(c.Request.Context(), uint(id))
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title": "Todo",
			"todo":  todo,
		})
	})

	r.GET("/todos/destory", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not invalid parameter"})
			return
		}
		todoRepo.Delete(c.Request.Context(), uint(id))
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.POST("/todos/update", func(c *gin.Context) {
		id, err := strconv.Atoi(c.PostForm("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not invalid parameter"})
			return
		}
		content := c.PostForm("content")
		var todo *models.Todo
		todo, _ = todoRepo.GetByID(c.Request.Context(), uint(id))
		todo.Content = content
		todoRepo.Update(c.Request.Context(), todo)
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.POST("/todos/create", func(c *gin.Context) {
		content := c.PostForm("content")
		todoRepo.Create(c, &models.Todo{Content: content})
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	fmt.Println("Database connection and setup successful")
	r.Run(":8000")
}
