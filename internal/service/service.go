package service

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"srep/internal/config"
	"strconv"
)

type Service struct {
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewService(cfg *config.Config, db *pgxpool.Pool) *Service {
	return &Service{cfg: cfg, db: db}
}

// Метод для создания персонажа
func (s *Service) CreateCharacter(name, species string, isForceUser bool, notes *string) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(),
		`INSERT INTO starwars_characters (name, species, is_force_user, notes) 
         VALUES ($1, $2, $3, $4) RETURNING id`,
		name, species, isForceUser, notes).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Метод для получения персонажа по ID
func (s *Service) GetCharacter(id int) (string, string, bool, *string, error) {
	var name, species string
	var isForceUser bool
	var notes *string
	err := s.db.QueryRow(context.Background(),
		`SELECT name, species, is_force_user, notes 
         FROM starwars_characters WHERE id=$1`, id).
		Scan(&name, &species, &isForceUser, &notes)
	if err != nil {
		return "", "", false, nil, err
	}
	return name, species, isForceUser, notes, nil
}

// Метод для обновления персонажа
func (s *Service) UpdateCharacter(id int, updates map[string]interface{}) error {
	// Проверяем, есть ли что обновлять
	if len(updates) == 0 {
		return nil
	}

	// Формируем SQL-запрос
	query := "UPDATE starwars_characters SET "
	args := []interface{}{}
	argIndex := 1

	for field, value := range updates {
		query += field + " = $" + strconv.Itoa(argIndex) + ", "
		args = append(args, value)
		argIndex++
	}

	// Убираем лишнюю запятую и пробел
	query = query[:len(query)-2]

	// Добавляем условие WHERE
	query += " WHERE id = $" + strconv.Itoa(argIndex)
	args = append(args, id)

	// Выполняем запрос
	_, err := s.db.Exec(context.Background(), query, args...)
	return err
}

// Метод для удаления персонажа
func (s *Service) DeleteCharacter(id int) error {
	_, err := s.db.Exec(context.Background(), "DELETE FROM starwars_characters WHERE id=$1", id)
	return err
}

// Метод для получения всех персонажей
func (s *Service) GetAllCharacters() ([]map[string]interface{}, error) {
	rows, err := s.db.Query(context.Background(),
		`SELECT id, name, species, is_force_user, notes 
         FROM starwars_characters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []map[string]interface{}
	for rows.Next() {
		var id int
		var name, species string
		var isForceUser bool
		var notes *string
		err = rows.Scan(&id, &name, &species, &isForceUser, &notes)
		if err != nil {
			return nil, err
		}

		character := map[string]interface{}{
			"id":            id,
			"name":          name,
			"species":       species,
			"is_force_user": isForceUser,
			"notes":         notes,
		}
		characters = append(characters, character)
	}

	return characters, nil
}
