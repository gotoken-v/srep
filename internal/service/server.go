package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"srep/internal/config"
	"srep/internal/validator"
)

func StartServer(cfg *config.Config, srv *Service) {
	app := fiber.New()

	// Структура запроса для создания персонажа
	type CharacterRequest struct {
		Name        string  `json:"name" validate:"required,min=3,max=50"`
		Species     string  `json:"species" validate:"required,min=3,max=50"`
		IsForceUser bool    `json:"is_force_user"`
		Notes       *string `json:"notes"`
	}

	// Обработчик для создания персонажа
	app.Post("/character", func(c *fiber.Ctx) error {
		var req CharacterRequest

		// Парсим тело запроса
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		// Валидируем запрос
		if err := validator.Validate(c.Context(), req); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Создаём персонажа
		id, err := srv.CreateCharacter(req.Name, req.Species, req.IsForceUser, req.Notes)
		if err != nil {
			return c.Status(500).SendString("Failed to create character")
		}

		return c.JSON(fiber.Map{"id": id})
	})

	// Обработчик для получения персонажа по ID
	app.Get("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		// Получаем персонажа
		name, species, isForceUser, notes, err := srv.GetCharacter(id)
		if err != nil {
			return c.Status(404).SendString("Character not found")
		}

		return c.JSON(fiber.Map{
			"id":            id,
			"name":          name,
			"species":       species,
			"is_force_user": isForceUser,
			"notes":         notes,
		})
	})

	// Обработчик для обновления персонажа
	app.Put("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		// Парсим тело запроса в карту
		var updates map[string]interface{}
		if err := c.BodyParser(&updates); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		// Обновляем данные персонажа
		err = srv.UpdateCharacter(id, updates)
		if err != nil {
			return c.Status(500).SendString("Failed to update character")
		}

		return c.SendString("Character updated successfully")
	})

	// Обработчик для удаления персонажа
	app.Delete("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		// Удаляем персонажа
		err = srv.DeleteCharacter(id)
		if err != nil {
			return c.Status(500).SendString("Failed to delete character")
		}

		return c.SendString("Character deleted successfully")
	})

	// Обработчик для получения всех персонажей
	app.Get("/characters", func(c *fiber.Ctx) error {
		// Получаем всех персонажей
		characters, err := srv.GetAllCharacters()
		if err != nil {
			return c.Status(500).SendString("Failed to retrieve characters")
		}

		return c.JSON(characters)
	})

	// Запуск сервера
	app.Listen(":8080")
}
