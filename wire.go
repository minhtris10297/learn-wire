//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"my-learning/golearn/handler"
	"my-learning/golearn/repository"
	"my-learning/golearn/service"

	"github.com/google/wire"
)

func Initialize(db *sql.DB) (*handler.MyHandler, error) {
	wire.Build(
		repository.NewMyRepository,
		service.NewMyService,
		handler.NewMyHandler,
	)

	return &handler.MyHandler{}, nil
}
