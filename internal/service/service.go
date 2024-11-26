package service

import (
	"github.com/gofiber/fiber/v2"
	"srep/internal/dto"
	"srep/internal/repo"
	"srep/pkg/validator"
	"strconv"
)

// Service реализует бизнес-логику приложения.
type Service struct {
	repo repo.RepositoryInterface
}

// NewService создаёт новый экземпляр Service с переданным репозиторием.
func NewService(repo repo.RepositoryInterface) *Service {
	return &Service{repo: repo}
}

// CreateCharacter обрабатывает запрос на создание персонажа.
func (s *Service) CreateCharacter(c *fiber.Ctx) error {
	var req dto.CharacterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := validator.Validate(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	character := repo.Character{
		Name:        "Unknown",
		Species:     "Unknown",
		IsForceUser: false,
		Notes:       req.Notes,
	}

	if req.Name != nil {
		character.Name = *req.Name
	}
	if req.Species != nil {
		character.Species = *req.Species
	}
	if req.IsForceUser != nil {
		character.IsForceUser = *req.IsForceUser
	}

	id, err := s.repo.CreateCharacter(c.Context(), character)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create character")
	}

	return c.JSON(fiber.Map{"id": id})
}

// UpdateCharacter обрабатывает запрос на обновление персонажа.
func (s *Service) UpdateCharacter(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	var req dto.CharacterRequest
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

	err = s.repo.UpdateCharacter(c.Context(), id, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update character")
	}

	return c.SendString("Character updated successfully")
}

// GetCharacter обрабатывает запрос на получение информации о персонаже.
func (s *Service) GetCharacter(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	character, err := s.repo.GetCharacter(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Character not found")
	}

	return c.JSON(character)
}

// DeleteCharacter обрабатывает запрос на удаление персонажа.
func (s *Service) DeleteCharacter(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	err = s.repo.DeleteCharacter(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete character")
	}

	return c.SendString("Character deleted successfully")
}

// GetAllCharacters обрабатывает запрос на получение всех персонажей.
func (s *Service) GetAllCharacters(c *fiber.Ctx) error {
	characters, err := s.repo.GetAllCharacters(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve characters")
	}

	return c.JSON(characters)
}
