package api

import (
	"github.com/gofiber/fiber/v2"
	"srep/internal/service"
)

func StartServer(svc service.ServiceInterface) {
	app := fiber.New()

	// Создание персонажа
	app.Post("/character", svc.CreateCharacter)

	// Обновление персонажа
	app.Put("/character/:id", svc.UpdateCharacter)

	// Получение персонажа
	app.Get("/character/:id", svc.GetCharacter)

	// Удаление персонажа
	app.Delete("/character/:id", svc.DeleteCharacter)

	// Получение всех персонажей
	app.Get("/characters", svc.GetAllCharacters)

	app.Listen(":8080")
}
