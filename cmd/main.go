package main

import (
	"github.com/joho/godotenv"
	"log"
	"srep/internal/config"
)

func main() {
	// Загружаем переменные из .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	config.LoadConfig()
}
