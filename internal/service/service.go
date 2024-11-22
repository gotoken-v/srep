package service

import (
	"context"
	"srep/internal/repo"
)

// Service реализует бизнес-логику приложения.
type Service struct {
	repo repo.RepositoryInterface
}

// NewService создаёт новый экземпляр Service с переданным репозиторием.
func NewService(repo repo.RepositoryInterface) *Service {
	return &Service{repo: repo}
}

// CreateCharacter создаёт нового персонажа.
func (s *Service) CreateCharacter(ctx context.Context, name, species string, isForceUser bool, notes *string) (int, error) {
	return s.repo.CreateCharacter(ctx, name, species, isForceUser, notes)
}

// GetCharacter получает информацию о персонаже по ID.
func (s *Service) GetCharacter(ctx context.Context, id int) (string, string, bool, *string, error) {
	return s.repo.GetCharacter(ctx, id)
}

// UpdateCharacter обновляет информацию о персонаже.
func (s *Service) UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error {
	return s.repo.UpdateCharacter(ctx, id, updates)
}

// DeleteCharacter удаляет персонажа по ID.
func (s *Service) DeleteCharacter(ctx context.Context, id int) error {
	return s.repo.DeleteCharacter(ctx, id)
}

// GetAllCharacters возвращает список всех персонажей.
func (s *Service) GetAllCharacters(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetAllCharacters(ctx)
}
