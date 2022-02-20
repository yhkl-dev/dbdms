package main

import (
	"dbdms/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

var databaseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Database",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"host": &graphql.Field{
				Type: graphql.String,
			},
			"port": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"schema": &graphql.Field{
				Type: graphql.String,
			},
			"comment": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"genre": &graphql.Field{
				Type: genreType,
			},
		},
	},
)

var databaseInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "databaseInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"host": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"port": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"username": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"schema": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"comment": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"genre_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

// graphql schema definition
var databaseFields = graphql.Fields{
	"getDatabaseByID": &graphql.Field{
		Type:        databaseType,
		Description: "Get database by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			fmt.Println(id)
			if ok {
				return app.models.DB.GetDatabaseByID(id)
			}
			return nil, nil
		},
	},
	"listDatabase": &graphql.Field{
		Name:        "listDatabase",
		Type:        graphql.NewList(databaseType),
		Description: "Get all databases",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return app.models.DB.ListDatabases()
		},
	},
	"searchDatabase": &graphql.Field{
		Type:        graphql.NewList(databaseType),
		Description: "Search database by name or host",
		Args: graphql.FieldConfigArgument{
			"nameContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"hostContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var databases []*models.Database
			name, okn := p.Args["nameContains"].(string)
			host, okh := p.Args["hostContains"].(string)
			result, err := app.models.DB.ListDatabases()
			if err != nil {
				return nil, err
			}
			for _, d := range result {
				if okn && strings.Contains(d.Name, name) {
					databases = append(databases, d)
				}
				if okh && strings.Contains(d.Host, host) {
					databases = append(databases, d)
				}
			}
			return databases, nil
		},
	},
}

var databaseMutationType = graphql.Fields{
	"testConnect": &graphql.Field{
		Type:        databaseType,
		Description: "test database connection",
		Args: graphql.FieldConfigArgument{
			"host": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"port": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"username": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"schema": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return nil, nil
		},
	},
	"createDatabase": &graphql.Field{
		Type: databaseType,
		Args: graphql.FieldConfigArgument{
			"databaseInputType": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(databaseInputType),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var params models.Database
			var x = p.Args["databaseInputType"].(map[string]interface{})
			err := mapstructure.Decode(x, &params)
			if err != nil {
				return nil, err
			}
			params.CreatedAt = time.Now()
			params.UpdatedAt = time.Now()
			if err = app.models.DB.CreateDatabase(params); err != nil {
				return nil, err
			}
			return params, nil
		},
	},
	"updateDatabase": &graphql.Field{
		Type: databaseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"host": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"port": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"username": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"schema": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"comment": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var params models.Database
			id := p.Args["id"]
			ID, _ := strconv.Atoi(id.(string))
			params.ID = ID
			params.Name = p.Args["name"].(string)
			params.Host = p.Args["host"].(string)
			params.Port = p.Args["port"].(int)
			params.Username = p.Args["username"].(string)
			params.Password = p.Args["password"].(string)
			params.Schema = p.Args["schema"].(string)
			params.Comment = p.Args["comment"].(string)
			params.UpdatedAt = time.Now()
			err := app.models.DB.UpdateDatabase(params)
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
	},
	"deleteDatabase": &graphql.Field{
		Type: databaseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["id"]
			ID, _ := strconv.Atoi(id.(string))
			err := app.models.DB.DeleteDatabase(ID)
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
	},
}
