package main

import (
	"github.com/FJC-OMUSUBI/go-handson/internals/config"
	"github.com/FJC-OMUSUBI/go-handson/internals/infrastructure/datasource"
	"github.com/FJC-OMUSUBI/go-handson/internals/interface/handler/api"
	"github.com/FJC-OMUSUBI/go-handson/internals/usecase"
)

func main() {
	postgresClient, cleanup, err := config.NewPostgres()
	if err != nil {
		panic(err)
	}
	transaction := config.NewTx(postgresClient)
	userRepository := datasource.NewUserRepository(postgresClient, transaction)
	tx := config.NewTx(postgresClient)
	userUsercase := usecase.NewUserUsecase(userRepository, tx)
	userHandler := api.NewUserHandler(userUsercase)
	engine := api.NewRouter(userHandler)
	engine.Run(":8080")
	defer cleanup()
}
