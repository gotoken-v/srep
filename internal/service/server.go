package service

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"srep/internal/config"
)

// StartServer запускает сервер на указанном порту.
func StartServer(cfg *config.Config) {
	app := fiber.New()

	// Обработчик для GET-запроса на корневой маршрут "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(cfg.SomeField)
	})

	// Запуск сервера на порту 8080
	log.Println("Сервер запущен на порту 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
