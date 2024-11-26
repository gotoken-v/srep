package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"srep/internal/config"
	"strconv"
)

type Repository struct {
	db *pgxpool.Pool
}

const (
	createCharacterQuery = `
		INSERT INTO starwars_characters (name, species, is_force_user, notes)
		VALUES ($1, $2, $3, $4) RETURNING id`
	getCharacterQuery = `
		SELECT name, species, is_force_user, notes
		FROM starwars_characters WHERE id=$1`
	updateCharacterQueryPrefix = "UPDATE starwars_characters SET "
	deleteCharacterQuery       = "DELETE FROM starwars_characters WHERE id=$1"
	getAllCharactersQuery      = `
		SELECT id, name, species, is_force_user, notes
		FROM starwars_characters`
)

// NewRepository создаёт новое подключение к базе данных
func NewRepository(cfg *config.Config) *Repository {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgreSQL.DBUser, cfg.PostgreSQL.DBPassword, cfg.PostgreSQL.DBHost, cfg.PostgreSQL.DBPort, cfg.PostgreSQL.DBName,
	)

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	return &Repository{db: db}
}

// Close закрывает соединение с базой данных
func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) CreateCharacter(ctx context.Context, name, species string, isForceUser bool, notes *string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, createCharacterQuery, name, species, isForceUser, notes).Scan(&id)
	return id, err
}

func (r *Repository) GetCharacter(ctx context.Context, id int) (string, string, bool, *string, error) {
	var name, species string
	var isForceUser bool
	var notes *string
	err := r.db.QueryRow(ctx, getCharacterQuery, id).Scan(&name, &species, &isForceUser, &notes)
	return name, species, isForceUser, notes, err
}

func (r *Repository) UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := updateCharacterQueryPrefix
	args := []interface{}{}
	argIndex := 1

	for field, value := range updates {
		query += field + " = $" + strconv.Itoa(argIndex) + ", "
		args = append(args, value)
		argIndex++
	}
	query = query[:len(query)-2]
	query += " WHERE id = $" + strconv.Itoa(argIndex)
	args = append(args, id)

	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *Repository) DeleteCharacter(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, deleteCharacterQuery, id)
	return err
}

func (r *Repository) GetAllCharacters(ctx context.Context) ([]map[string]interface{}, error) {
	rows, err := r.db.Query(ctx, getAllCharactersQuery)
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
