package repositories

import (
	"context"
	"gin_todo_app_clean_archi/src/domain/models"
)

type TodoRepository interface {
	GetByID(ctx context.Context, id uint) (*models.Todo, error)
	Create(ctx context.Context, todo *models.Todo) error
	Update(ctx context.Context, todo *models.Todo) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*models.Todo, error)
}
