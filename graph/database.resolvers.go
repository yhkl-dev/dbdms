package graph

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/yhkl-dev/dbdms/graph/model"
)

func (r *mutationResolver) UpdateDatabase(ctx context.Context, input model.UpdateDatabaseInput) (*model.Database, error) {
	stmt := `
		update t_databases set
			name = $1,
			host = $2,
			port = $3,
			schema = $4,
			username = $5,
			password = $6,
			comment = $7,
			genre_id = $8
		where 
			id = $9
	`
	_, err := r.DB.ExecContext(ctx, stmt, input.Name, input.Host, input.Port, input.Schema, input.Username, input.Password, input.Comment, input.Genre.ID, input.ID)
	if err != nil {
		return nil, err
	}
	queryString := `
			select
			t1.id,
			t1.name,
			t1.host,
			t1.port,
			t1.username,
			t1.schema,
			t1.comment,
			t1.created_at,
			t1.updated_at,
			t2.id,
			t2.genre_name,
			t3.id,
			t3.username
		from
			t_databases t1,
			t_genre t2,
			t_user t3
		where
			t1.genre_id = t2.id
			and t1.user_id = t3.id
			and t1.id = $1
		`
	rows := r.DB.QueryRowContext(ctx, queryString, input.ID)

	var database model.Database
	var genre model.Genre
	var user model.User
	err = rows.Scan(
		&database.ID,
		&database.Name,
		&database.Host,
		&database.Port,
		&database.Username,
		&database.Schema,
		&database.Comment,
		&database.CreateAt,
		&database.UpdateAt,
		&genre.ID,
		&genre.Name,
		&user.ID,
		&user.Name,
	)
	database.Genre = &genre
	database.User = &user
	if err != nil {
		return nil, err
	}
	return &database, nil
}

func (r *mutationResolver) DetelDatabase(ctx context.Context, input string) (string, error) {
	id, _ := strconv.Atoi(input)
	stmt := `
		delete from t_databases where 
			id = $1
	`
	_, err := r.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return "error", err
	}
	return "ok", nil
}

func (r *mutationResolver) CreateDatabase(ctx context.Context, input *model.NewDatabaseInput) (*model.Database, error) {
	stmt := `
		insert into t_databases (
			name, host, port, username, password, schema, comment, created_at, updated_at, genre_id, user_id
		)
		values(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
		returning id
	`
	var insertID int
	err := r.DB.QueryRow(stmt,
		input.Name, input.Host, input.Port, input.Username, input.Password,
		input.Schema, input.Schema, time.Now(), time.Now(), input.Genre.ID, input.User.ID).
		Scan(&insertID)
	if err != nil {
		return nil, err
	}

	queryString := `
			select
			t1.id,
			t1.name,
			t1.host,
			t1.port,
			t1.username,
			t1.schema,
			t1.comment,
			t1.created_at,
			t1.updated_at,
			t2.id,
			t2.genre_name,
			t3.id,
			t3.username
		from
			t_databases t1,
			t_genre t2,
			t_user t3
		where
			t1.genre_id = t2.id
			and t1.user_id = t3.id
			and t1.id = $1
		`
	rows := r.DB.QueryRowContext(ctx, queryString, insertID)

	var database model.Database
	var genre model.Genre
	var user model.User
	err = rows.Scan(
		&database.ID,
		&database.Name,
		&database.Host,
		&database.Port,
		&database.Username,
		&database.Schema,
		&database.Comment,
		&database.CreateAt,
		&database.UpdateAt,
		&genre.ID,
		&genre.Name,
		&user.ID,
		&user.Name,
	)
	database.Genre = &genre
	database.User = &user
	if err != nil {
		return nil, err
	}
	return &database, nil
}

func (r *queryResolver) QueryDatabases(ctx context.Context, input *model.Page) ([]*model.Database, error) {
	limitString := ""
	if input != nil {
		limitString = fmt.Sprintf("limit %d offset %d", *input.PageSize, *input.Page**input.PageSize)
	}
	queryString := fmt.Sprintf(`
			select
			t1.id,
			t1.name,
			t1.host,
			t1.port,
			t1.username,
			t1.schema,
			t1.comment,
			t1.created_at,
			t1.updated_at,
			t2.id,
			t2.genre_name,
			t3.id,
			t3.username
		from
			t_databases t1,
			t_genre t2,
			t_user t3
		where
			t1.genre_id = t2.id
			and t1.user_id = t3.id
			%s
		`, limitString)
	rows, err := r.DB.QueryContext(ctx, queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []*model.Database
	for rows.Next() {
		var database model.Database
		var genre model.Genre
		var user model.User
		err := rows.Scan(
			&database.ID,
			&database.Name,
			&database.Host,
			&database.Port,
			&database.Username,
			&database.Schema,
			&database.Comment,
			&database.CreateAt,
			&database.UpdateAt,
			&genre.ID,
			&genre.Name,
			&user.ID,
			&user.Name,
		)
		database.Genre = &genre
		database.User = &user
		if err != nil {
			return nil, err
		}
		databases = append(databases, &database)
	}
	return databases, nil
}

func (r *queryResolver) QueryDatabaseByID(ctx context.Context, input string) (*model.Database, error) {
	queryString := `
			select
			t1.id,
			t1.name,
			t1.host,
			t1.port,
			t1.username,
			t1.schema,
			t1.comment,
			t1.created_at,
			t1.updated_at,
			t2.id,
			t2.genre_name,
			t3.id,
			t3.username
		from
			t_databases t1,
			t_genre t2,
			t_user t3
		where
			t1.genre_id = t2.id
			and t1.user_id = t3.id
			and t1.id = $1
	`
	inputID, _ := strconv.Atoi(input)
	rows := r.DB.QueryRowContext(ctx, queryString, inputID)
	var database model.Database
	var genre model.Genre
	var user model.User
	err := rows.Scan(
		&database.ID,
		&database.Name,
		&database.Host,
		&database.Port,
		&database.Username,
		&database.Schema,
		&database.Comment,
		&database.CreateAt,
		&database.UpdateAt,
		&genre.ID,
		&genre.Name,
		&user.ID,
		&user.Name,
	)
	database.Genre = &genre
	database.User = &user
	if err != nil {
		return nil, err
	}
	return &database, nil
}
