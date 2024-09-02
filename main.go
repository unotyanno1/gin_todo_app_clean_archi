package main

import (
	"log"
	"gin_todo_app_clean_archi/src/infra/database/repositories"
	"gin_todo_app_clean_archi/src/infra/database"
	"gin_todo_app_clean_archi/src/usecase/services"
	"gin_todo_app_clean_archi/src/interface/controllers"
	"gin_todo_app_clean_archi/src/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	todoRepo := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	r := gin.Default()
	r = router.SetupRouterTodo(r, todoController)
	r = router.SetupRouterPage(r, todoController)

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
