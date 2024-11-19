package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"srep/internal/validator"
)

func StartServer(srv *Service) {
	app := fiber.New()

	// Универсальная структура для создания и обновления персонажа
	type CharacterRequest struct {
		Name        *string `json:"name" validate:"omitempty,min=3,max=50,name"`
		Species     *string `json:"species" validate:"omitempty,min=3,max=50,species"`
		IsForceUser *bool   `json:"is_force_user" validate:"omitempty,force_user"`
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

		// Устанавливаем значения по умолчанию для необязательных полей
		name := "Unknown"
		if req.Name != nil {
			name = *req.Name
		}

		species := "Unknown Species"
		if req.Species != nil {
			species = *req.Species
		}

		isForceUser := false
		if req.IsForceUser != nil {
			isForceUser = *req.IsForceUser
		}

		// Создаём персонажа
		id, err := srv.CreateCharacter(name, species, isForceUser, req.Notes)
		if err != nil {
			return c.Status(500).SendString("Failed to create character")
		}

		return c.JSON(fiber.Map{"id": id})
	})

	// Обработчик для обновления персонажа
	app.Put("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		// Парсим тело запроса
		var req CharacterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		// Валидируем только переданные поля
		if err := validator.Validate(c.Context(), req); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Преобразуем структуру в карту для обновления
		updates := map[string]interface{}{}
		if req.Name != nil {
			updates["name"] = *req.Name
		}
		if req.Species != nil {
			updates["species"] = *req.Species
		}
		if req.IsForceUser != nil {
			updates["is_force_user"] = *req.IsForceUser
		}
		if req.Notes != nil {
			updates["notes"] = *req.Notes
		}

		// Обновляем данные персонажа
		err = srv.UpdateCharacter(id, updates)
		if err != nil {
			return c.Status(500).SendString("Failed to update character")
		}

		return c.SendString("Character updated successfully")
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
