package main

import (
	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Query",
	Fields: databaseFields,
})
