package models

import (
	"context"
	"time"
)

type Document struct {
	ID         int       `json:"id"`
	DatabaseID int       `json:"database_id"`
	Content    string    `json:"contect"`
	CreateAt   time.Time `json:"create_at"`
}

func (m *DBModel) ListGenres() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from t_genre`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre
	for rows.Next() {
		var genre Genre
		err := rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &genre)
	}
	return genres, nil
}

func (m *DBModel) CreateGenre(data Genre) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `
		insert into t_genre (genre_name, create_at, update_at)
		values ($1, $2, $3)
	`
	_, err := m.DB.ExecContext(ctx, stmt, data.GenreName, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
