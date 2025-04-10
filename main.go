package main

import (
	"log"
	"my-todo-app/config"
	"my-todo-app/internal/api"
	"my-todo-app/internal/repo"
	"my-todo-app/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация репозитория
	repo := repo.NewRepository()

	// Создание сервиса
	service := service.NewService(repo)

	// Настройка роутов
	app := api.NewRouter(service)

	// Запуск сервера
	log.Printf("Starting server on %s", cfg.Port)
	if err := app.Listen(cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
