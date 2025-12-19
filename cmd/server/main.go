package main

import "log"

func main() {
	// 1. Инициализация БД
	database.InitDB()

	// 2. Репозиторий и сервис задач
	repo := task.NewRepository(database.DB)
	svc := task.NewService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
