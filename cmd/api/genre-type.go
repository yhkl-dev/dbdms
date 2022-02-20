package main

import (
	"dbdms/models"
	"time"

	"github.com/graphql-go/graphql"
)

var genreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Genre",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"genre_name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var genreFields = graphql.Fields{
	"listGenre": &graphql.Field{
		Name:        "listGenre",
		Type:        graphql.NewList(genreType),
		Description: "Get all Genres",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return app.models.DB.ListGenres()
		},
	},
}

var genreMutationType = graphql.Fields{
	"createGenre": &graphql.Field{
		Type:        genreType,
		Description: "create genre",
		Args: graphql.FieldConfigArgument{
			"genre_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var params models.Genre
			genreName := p.Args["name"].(string)
			params.GenreName = genreName
			params.CreatedAt = time.Now()
			params.UpdatedAt = time.Now()

			if err := app.models.DB.CreateGenre(params); err != nil {
				return nil, err
			}
			return params, nil
		},
	},
}
