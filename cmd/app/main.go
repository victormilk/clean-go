package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"todoList/configs"
	"todoList/internal/infra/web/webserver"
	db "todoList/sql/sqlc"
)

func main() {
	conf := configs.LoadConfig()
	ctx := context.Background()

	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		conf.DBUser, conf.DBPassword, conf.DBName, conf.SSLMode, conf.DBHost, conf.DBPort)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	todoHandler := NewTodoHandler(queries)
	webServer := webserver.NewWebServer("8080")
	todoHandler.RegisterRoutes(webServer)

	log.Println("Starting server on port 8080")
	webServer.Start()
}
