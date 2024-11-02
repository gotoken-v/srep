package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"srep/internal/config"
	"srep/internal/service"
)

func main() {
	// Загружаем переменные из .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	srv := service.NewService(cfg)
	field := srv.GetConfigField()
	fmt.Println(field)
}
