// Package service provides the business logic for the todo endpoint.
package service

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
)

// Todo is the service for the todo endpoint.
type Todo interface {
	Create(todo model.Todo) (*model.Todo, error)
	Update(todo model.Todo) (*model.Todo, error)
	Delete(id int) error
	Find(id int) (*model.Todo, error)
	FindAll(query *model.TodoQuery) ([]*model.Todo, error)
}

type todo struct {
	todoRepository repository.Todo
}

// NewTodo creates a new Todo service.
func NewTodo(r repository.Todo) Todo {
	return &todo{r}
}

func (t *todo) Create(todoCreate model.Todo) (*model.Todo, error) {
	td := &model.Todo{
		Task:     todoCreate.Task,
		Priority: todoCreate.Priority,
	}
	if err := t.todoRepository.Create(td); err != nil {
		return nil, err
	}
	return td, nil
}

func (t *todo) Update(todoUpdate model.Todo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:       todoUpdate.ID,
		Task:     todoUpdate.Task,
		Priority: todoUpdate.Priority,
		Status:   todoUpdate.Status,
	}
	// 現在の値を取得
	currentTodo, err := t.Find(todo.ID)
	if err != nil {
		return nil, err
	}
	// 空文字列の場合、現在の値を使用
	if todo.Task == "" {
		todo.Task = currentTodo.Task
	}
	if todo.Status == "" {
		todo.Status = currentTodo.Status
	}

	if todo.Priority <= 0 {
		todo.Priority = currentTodo.Priority
	}

	if err := t.todoRepository.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) Delete(id int) error {
	if err := t.todoRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *todo) Find(id int) (*model.Todo, error) {
	todo, err := t.todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) FindAll(query *model.TodoQuery) ([]*model.Todo, error) {
	todo, err := t.todoRepository.FindAll(query)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
