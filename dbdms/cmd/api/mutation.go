package main

import (
	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: databaseMutationType,
})
