package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"srep/internal/config"
	"srep/internal/service"
)

func main() {

	// Загружаем переменные окружения из файла .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Формируем строку подключения к базе данных
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Подключаемся к базе данных через пул соединений
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Создаем сервис и передаем в него конфигурацию и подключение к базе данных
	srv := service.NewService(cfg, db)

	// Запускаем сервер и передаем в него конфигурацию и созданный сервис
	service.StartServer(cfg, srv)
}
