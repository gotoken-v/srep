package service

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
)

// RepositoryInterface определяет методы для взаимодействия с базой данных.
type RepositoryInterface interface {
	CreateCharacter(ctx context.Context, name, species string, isForceUser bool, notes *string) (int, error)
	GetCharacter(ctx context.Context, id int) (string, string, bool, *string, error)
	UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error
	DeleteCharacter(ctx context.Context, id int) error
	GetAllCharacters(ctx context.Context) ([]map[string]interface{}, error)
}

// Repository реализует методы взаимодействия с реальной базой данных.
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository создаёт новый экземпляр Repository.
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateCharacter(ctx context.Context, name string, species string, isForceUser bool, notes *string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, `
        INSERT INTO starwars_characters (name, species, is_force_user, notes)
        VALUES ($1, $2, $3, $4) RETURNING id`, name, species, isForceUser, notes).Scan(&id)
	return id, err
}

func (r *Repository) GetCharacter(ctx context.Context, id int) (string, string, bool, *string, error) {
	var name, species string
	var isForceUser bool
	var notes *string
	err := r.db.QueryRow(ctx, `
        SELECT name, species, is_force_user, notes
        FROM starwars_characters WHERE id=$1`, id).Scan(&name, &species, &isForceUser, &notes)
	return name, species, isForceUser, notes, err
}

func (r *Repository) UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE starwars_characters SET "
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
	_, err := r.db.Exec(ctx, "DELETE FROM starwars_characters WHERE id=$1", id)
	return err
}

func (r *Repository) GetAllCharacters(ctx context.Context) ([]map[string]interface{}, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, name, species, is_force_user, notes
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
