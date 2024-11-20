package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"srep/internal/config"
	"srep/internal/service"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Формируем строку подключения к базе данных
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	// Подключаемся к базе данных через пул соединений
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Создаём реальный репозиторий
	repo := service.NewRepository(db)

	// Создаём сервис
	srv := service.NewService(repo)

	// Запускаем сервер
	service.StartServer(srv)
}
