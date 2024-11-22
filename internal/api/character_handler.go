package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"srep/internal/service"
	"srep/internal/validator"
)

func StartServer(svc service.ServiceInterface) {
	app := fiber.New()

	type CharacterRequest struct {
		Name        *string `json:"name" validate:"omitempty,min=3,max=50,name"`
		Species     *string `json:"species" validate:"omitempty,min=3,max=50,species"`
		IsForceUser *bool   `json:"is_force_user" validate:"omitempty,force_user"`
		Notes       *string `json:"notes"`
	}

	// Создание персонажа
	app.Post("/character", func(c *fiber.Ctx) error {
		var req CharacterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if err := validator.Validate(c.Context(), req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		name := "Unknown"
		if req.Name != nil {
			name = *req.Name
		}
		species := "Unknown"
		if req.Species != nil {
			species = *req.Species
		}
		isForceUser := false
		if req.IsForceUser != nil {
			isForceUser = *req.IsForceUser
		}

		id, err := svc.CreateCharacter(c.Context(), name, species, isForceUser, req.Notes)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create character")
		}

		return c.JSON(fiber.Map{"id": id})
	})

	// Обновление персонажа
	app.Put("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		var req CharacterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if err := validator.Validate(c.Context(), req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

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

		err = svc.UpdateCharacter(c.Context(), id, updates)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update character")
		}

		return c.SendString("Character updated successfully")
	})

	// Получение персонажа
	app.Get("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		name, species, isForceUser, notes, err := svc.GetCharacter(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Character not found")
		}

		return c.JSON(fiber.Map{
			"id":            id,
			"name":          name,
			"species":       species,
			"is_force_user": isForceUser,
			"notes":         notes,
		})
	})

	// Удаление персонажа
	app.Delete("/character/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		err = svc.DeleteCharacter(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete character")
		}

		return c.SendString("Character deleted successfully")
	})

	// Получение всех персонажей
	app.Get("/characters", func(c *fiber.Ctx) error {
		characters, err := svc.GetAllCharacters(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve characters")
		}

		return c.JSON(characters)
	})

	app.Listen(":8080")
}
