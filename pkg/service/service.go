package service

import (
	todo "WebApplication"
	"WebApplication/pkg/repository"
)

/**
 * Authorization, TodoList, TodoItem main Service
 */

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
