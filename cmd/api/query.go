package main

import (
	"github.com/graphql-go/graphql"
)

var graphqlFields = graphql.Fields{}

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Query",
	Fields: merge(graphqlFields, databaseFields, genreFields),
})
