package usecase_services_test

import (
    "context"
	"testing"
	_"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gin_todo_app_clean_archi/src/domain/models"
	"gin_todo_app_clean_archi/src/usecase/services"
)

type MockTodoRepository struct {
    mock.Mock
}

func (m *MockTodoRepository) Create(ctx context.Context, todo *models.Todo) error {
    args := m.Called(ctx, todo)
    return args.Error(0)
}

func (m *MockTodoRepository) GetByID(ctx context.Context, id uint) (*models.Todo, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(ctx context.Context, todo *models.Todo) error {
    args := m.Called(ctx, todo)
    return args.Error(0)
}

func (m *MockTodoRepository) Delete(ctx context.Context, id uint) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}

func (m *MockTodoRepository) List(ctx context.Context) ([]*models.Todo, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*models.Todo), args.Error(1)
}

func TestCreateTodo(t *testing.T) {
    setup()
    todo := &models.Todo{Content: "test"}
    mockRepo.On("Create", mock.Anything, todo).Return(nil)
    err := service.CreateTodo(ctx, "test")
    assert.Nil(t, err)
    mockRepo.AssertExceptations(t)
}

func TestUpdateTodo(t *testing.T) {
    setup()
    todo := &models.Todo{Content: "test"}
    mockRepo.On("Update", mock.Anything, todo).Return(nil)
    err := service.UpdateTodo(ctx, todo)
    assert.Nil(t, err)
    mockRepo.AssertExceptations(t)
}

func TestDeleteTodo(t *testing.T){
    setup()
    mockRepo.On("Delete", mock.Anything, uint(1)).Return(nil)
    err := service.DeleteTodo(ctx, 1)
    assert.Nil(t, err)
    mockRepo.AssertExceptations(t)
}

func TestListTodos(t *testing.T){
    setup()
    mockTodos := []*models.Todo{{Content: "test"}, {Content: "clean architecture"}}
    mockRepo.On("List", mock.Anything).Return(mockTodos, nil)
    todos, err := service.ListTodos(ctx)
    assert.Nil(t, err)
    assert.Equal(t, mockTodos, todos)
    mockRepo.AssertExceptations(t)
}

var (
    mockRepo *MockTodoRepository
    service *services.TodoService
    ctx context.Context
)

func setup() {
    mockRepo = new(MockTodoRepository)
    service = services.NewTodoService(mockRepo)
    ctx = context.Background()
}