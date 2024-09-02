package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_todo_app_clean_archi/src/interface/controllers"
)

func SetupRouterPage(r *gin.Engine, todoController *controllers.TodoController) *gin.Engine {
    r.LoadHTMLGlob("src/infra/http/public/*")

    r.GET("/index", func(c *gin.Context) {
		todos, err := todoController.ListTodos(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve todos"})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "やることリスト(クリーンアーキテクチャー版)",
			"todos": todos,
		})
	})

	r.GET("/todos/edit", func(c *gin.Context) {
		todo, err := todoController.GetTodoByID(c)
		if err != nil {
			return
		}
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title": "Todoの編集",
			"todo":  todo,
		})
	})

	return r
}