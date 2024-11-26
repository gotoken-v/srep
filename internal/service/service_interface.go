package service

import "github.com/gofiber/fiber/v2"

// ServiceInterface определяет методы бизнес-логики приложения.
type ServiceInterface interface {
	CreateCharacter(c *fiber.Ctx) error
	UpdateCharacter(c *fiber.Ctx) error
	GetCharacter(c *fiber.Ctx) error
	DeleteCharacter(c *fiber.Ctx) error
	GetAllCharacters(c *fiber.Ctx) error
}
