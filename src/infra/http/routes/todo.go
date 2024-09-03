package router

import (
    "github.com/gin-gonic/gin"
    "gin_todo_app_clean_archi/src/interface/controllers"
)

func SetupRouterTodo(r *gin.Engine, todoController *controllers.TodoController) *gin.Engine {
    r.POST("/todos/create", todoController.CreateTodo)
    r.POST("/todos/update", todoController.UpdateTodo)
    r.GET("/todos/destory", todoController.DeleteTodo)

    return r
}