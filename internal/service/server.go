package service

import (
	"github.com/gofiber/fiber/v2"
	"srep/internal/config"
	"strconv"
)

func StartServer(cfg *config.Config, srv *Service) {
	app := fiber.New()

	// Обработчик для создания персонажа
	app.Post("/character", func(c *fiber.Ctx) error {
		type request struct {
			Name    string `json:"name"`
			Species string `json:"species"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		id, err := srv.CreateCharacter(req.Name, req.Species)
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

		name, species, err := srv.GetCharacter(id)
		if err != nil {
			return c.Status(404).SendString("Character not found")
		}

		return c.JSON(fiber.Map{"id": id, "name": name, "species": species})
	})

	// Обработчик для обновления персонажа
	app.Put("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		type request struct {
			Name    string `json:"name"`
			Species string `json:"species"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		err = srv.UpdateCharacter(id, req.Name, req.Species)
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

		err = srv.DeleteCharacter(id)
		if err != nil {
			return c.Status(500).SendString("Failed to delete character")
		}

		return c.SendString("Character deleted successfully")
	})

	// Обработчик для получения всех персонажей
	app.Get("/characters", func(c *fiber.Ctx) error {
		characters, err := srv.GetAllCharacters()
		if err != nil {
			return c.Status(500).SendString("Failed to retrieve characters")
		}

		return c.JSON(characters)
	})

	// Запуск сервера
	app.Listen(":8080")
}
