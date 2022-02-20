package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewDBModel
func NewDBModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Database struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Port      int       `json:"port"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Schema    string    `json:"schema"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GenreName string    `json:"genre_name"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type DatabaseGenres struct {
	ID         int
	DatabaseID int       `json:"-"`
	GenreID    int       `json:"-"`
	Genre      Genre     `json:"genre"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type DatabaseContent struct {
	ID         int       `json:"id"`
	VersionID  string    `json:"version_id"`
	DatabaseID int       `json:"database_id"`
	Document   string    `json:"ducoment"`
	CreatedAt  time.Time `json:"created_at"`
}

type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
