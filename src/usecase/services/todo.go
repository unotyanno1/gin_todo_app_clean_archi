package services

import (
    "context"
    "gin_todo_app_clean_archi/src/domain/models"
    "gin_todo_app_clean_archi/src/domain/repositories"
)

type TodoService struct {
    repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
    return &TodoService{
        repo: repo,
    }
}

func (s *TodoService) CreateTodo(ctx context.Context, content string) error {
    todo := &models.Todo{Content: content}
    return s.repo.Create(ctx, todo)
}

func (s *TodoService) GetTodoByID(ctx context.Context, id uint) (*models.Todo, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *TodoService) UpdateTodo(ctx context.Context, todo *models.Todo) error {
    if err := todo.Validate(); err != nil {
        return err
    }
    return s.repo.Update(ctx, todo)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *TodoService) ListTodos(ctx context.Context) ([]*models.Todo, error) {
    return s.repo.List(ctx)
}