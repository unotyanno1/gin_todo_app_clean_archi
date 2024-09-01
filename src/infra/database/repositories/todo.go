package repository

import (
	"context"
	"gin_todo_app_clean_archi/src/domain/models"
	"gin_todo_app_clean_archi/src/domain/repositories"

	"gorm.io/gorm"

)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repositories.TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetByID(ctx context.Context, id uint) (*models.Todo, error) {
	var todo models.Todo
	result := r.DB.First(&todo, id)
	return &todo, result.Error
}

func (r *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) Update(ctx context.Context, todo *models.Todo) error {
	return r.DB.WithContext(ctx).Save(todo).Error
}

func (r *TodoRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.Delete(&models.Todo{}, id).Error
}

func (r *TodoRepository) List(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	result := r.DB.Find(&todos)
	return todos, result.Error
}

