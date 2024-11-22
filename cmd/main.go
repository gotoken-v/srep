package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"srep/internal/api"
	"srep/internal/config"
	"srep/internal/repo"
	"srep/internal/service"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключение к базе данных
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Инициализация компонентов
	repository := repo.NewRepository(db)
	service := service.NewService(repository)

	// Запуск HTTP-сервера
	api.StartServer(service)
}
