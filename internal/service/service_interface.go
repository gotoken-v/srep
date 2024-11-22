package service

import "context"

// ServiceInterface определяет методы бизнес-логики приложения.
type ServiceInterface interface {
	// Создание нового персонажа
	CreateCharacter(ctx context.Context, name, species string, isForceUser bool, notes *string) (int, error)

	// Получение персонажа по ID
	GetCharacter(ctx context.Context, id int) (string, string, bool, *string, error)

	// Обновление персонажа
	UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error

	// Удаление персонажа
	DeleteCharacter(ctx context.Context, id int) error

	// Получение списка всех персонажей
	GetAllCharacters(ctx context.Context) ([]map[string]interface{}, error)
}
