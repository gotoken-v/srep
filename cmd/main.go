package main

import (
	"srep/internal/api"
	"srep/internal/config"
	"srep/internal/repo"
	"srep/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Ошибка загрузки конфигурации: " + err.Error())
	}

	// Создание репозитория
	repository := repo.NewRepository(cfg)
	defer repository.Close()

	// Создание сервиса
	service := service.NewService(repository)

	// Запуск HTTP-сервера
	api.StartServer(service)
}
