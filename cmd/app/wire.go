//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"todoList/internal/entity"
	"todoList/internal/infra/database"
	"todoList/internal/infra/web"
	db "todoList/sql/sqlc"
)

var setTodoRepositoryDependency = wire.NewSet(
	database.NewTodoRepository,
	wire.Bind(new(entity.TodoRepositoryInterface), new(*database.TodoRepository)),
)

func NewTodoHandler(queries *db.Queries) *web.TodoHandler {
	wire.Build(setTodoRepositoryDependency, web.NewTodoHandler)
	return &web.TodoHandler{}
}
