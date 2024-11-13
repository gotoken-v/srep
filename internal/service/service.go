package service

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"srep/internal/config"
)

type Service struct {
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewService(cfg *config.Config, db *pgxpool.Pool) *Service {
	return &Service{cfg: cfg, db: db}
}

// Метод для создания персонажа
func (s *Service) CreateCharacter(name, species string) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(),
		"INSERT INTO starwars_characters (name, species) VALUES ($1, $2) RETURNING id",
		name, species).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Метод для получения персонажа по ID
func (s *Service) GetCharacter(id int) (string, string, error) {
	var name, species string
	err := s.db.QueryRow(context.Background(),
		"SELECT name, species FROM starwars_characters WHERE id=$1", id).Scan(&name, &species)
	if err != nil {
		return "", "", err
	}
	return name, species, nil
}

// Метод для обновления персонажа
func (s *Service) UpdateCharacter(id int, name, species string) error {
	_, err := s.db.Exec(context.Background(),
		"UPDATE starwars_characters SET name=$1, species=$2 WHERE id=$3", name, species, id)
	return err
}

// Метод для удаления персонажа
func (s *Service) DeleteCharacter(id int) error {
	_, err := s.db.Exec(context.Background(), "DELETE FROM starwars_characters WHERE id=$1", id)
	return err
}

// Метод для получения всех персонажей
func (s *Service) GetAllCharacters() ([]map[string]interface{}, error) {
	rows, err := s.db.Query(context.Background(), "SELECT id, name, species FROM starwars_characters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []map[string]interface{}
	for rows.Next() {
		var id int
		var name, species string
		err = rows.Scan(&id, &name, &species)
		if err != nil {
			return nil, err
		}

		character := map[string]interface{}{
			"id":      id,
			"name":    name,
			"species": species,
		}
		characters = append(characters, character)
	}

	return characters, nil
}
