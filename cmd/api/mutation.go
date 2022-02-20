package main

import (
	"github.com/graphql-go/graphql"
)

var graphqlMutationFields = graphql.Fields{}

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: merge(graphqlMutationFields, databaseMutationType),
})
