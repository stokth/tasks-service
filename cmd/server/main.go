package main

import (
	"log"

	"github.com/stokth/tasks-service/internal/database"
	"github.com/stokth/tasks-service/internal/task"
	"github.com/stokth/tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 2. Репозиторий и сервис задач
	repo := task.NewRepository(db)
	svc := task.NewService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := grpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	if err := grpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
