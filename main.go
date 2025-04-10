package main

import (
	"log"
	"my-todo-app/config"
	"my-todo-app/internal/api"
	"my-todo-app/internal/repo"
	"my-todo-app/internal/service"

	"github.com/joho/godotenv" // <--- Добавленный импорт
)

func main() {
	// Загрузка переменных из .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Ошибка загрузки .env: %v", err)
	}

	// Загрузка конфигурации через envconfig
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
