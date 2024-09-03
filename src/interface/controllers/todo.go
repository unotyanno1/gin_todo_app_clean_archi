package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "gin_todo_app_clean_archi/src/usecase/services"
    "gin_todo_app_clean_archi/src/domain/models"
)

type TodoController struct {
    service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
    return &TodoController{
        service: service,
    }
}

func (tc *TodoController) GetTodoByID(c *gin.Context) (*models.Todo, error) {
    id, err := strconv.Atoi(c.Query("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "inavlid id"})
        return nil, err
    }
    return tc.service.GetTodoByID(c.Request.Context(), uint(id))
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
    content := c.PostForm("content")
    if content == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "content is required"})
        return
    }
    err := tc.service.CreateTodo(c.Request.Context(), content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create todo"})
        return
    }
    c.Status(http.StatusCreated)
    c.Redirect(http.StatusMovedPermanently, "/index")
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
    idStr := c.PostForm("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    content := c.PostForm("content")
    todo, err := tc.service.GetTodoByID(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Content is not found"})
        return
    }
    todo.Content = content
    if err := tc.service.UpdateTodo(c.Request.Context(), todo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
    c.Redirect(http.StatusSeeOther, "/index")
}

func (tc *TodoController) DeleteTodo(c *gin.Context) {
    idStr, ok := c.GetQuery("id")
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID query parameter is required"})
        return
    }
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    err = tc.service.DeleteTodo(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }
    c.Status(http.StatusOK)
    c.Redirect(http.StatusSeeOther, "/index")
}

func (tc *TodoController) ListTodos(c *gin.Context) ([]*models.Todo, error) {
    return tc.service.ListTodos(c.Request.Context())
}

