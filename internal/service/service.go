package service

import "context"

// Service представляет бизнес-логику приложения.
type Service struct {
	repo RepositoryInterface
}

// NewService создаёт новый экземпляр Service с переданным репозиторием.
func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

// CreateCharacter создаёт нового персонажа.
func (s *Service) CreateCharacter(name, species string, isForceUser bool, notes *string) (int, error) {
	return s.repo.CreateCharacter(context.Background(), name, species, isForceUser, notes)
}

// GetCharacter получает информацию о персонаже по ID.
func (s *Service) GetCharacter(id int) (string, string, bool, *string, error) {
	return s.repo.GetCharacter(context.Background(), id)
}

// UpdateCharacter обновляет информацию о персонаже.
func (s *Service) UpdateCharacter(id int, updates map[string]interface{}) error {
	return s.repo.UpdateCharacter(context.Background(), id, updates)
}

// DeleteCharacter удаляет персонажа по ID.
func (s *Service) DeleteCharacter(id int) error {
	return s.repo.DeleteCharacter(context.Background(), id)
}

// GetAllCharacters возвращает список всех персонажей.
func (s *Service) GetAllCharacters() ([]map[string]interface{}, error) {
	return s.repo.GetAllCharacters(context.Background())
}
