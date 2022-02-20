package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

func (app *application) handle(ctx context.Context, h *handler.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var opts RequestOptions
		json.Unmarshal(bodyBytes, &opts)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		h.ContextHandler(ctx, w, r)
	}
}

func RegisterSchema() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    Query,
			Mutation: Mutation,
		})
	if err != nil {
		panic(fmt.Sprintf("schema init fail %s", err.Error()))
	}
	return &schema
}

func Register() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   RegisterSchema(),
		Pretty:   true,
		GraphiQL: true,
	})
}

// func (app *application) ListDatabaseGraphQL(w http.ResponseWriter, r *http.Request) {
// 	databases, _ = app.models.DB.ListDatabases()
// 	q, _ := io.ReadAll(r.Body)
// 	query := string(q)
// 	log.Println(query)

// 	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
// 	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
// 	schema, err := graphql.NewSchema(schemaConfig)
// 	if err != nil {
// 		app.errorJSON(w, errors.New("failed to creat schema"))
// 		log.Println(err)
// 		return
// 	}

// 	params := graphql.Params{Schema: schema, RequestString: query}
// 	response := graphql.Do(params)
// 	if len(response.Errors) > 0 {
// 		app.errorJSON(w, fmt.Errorf("failed: %+v", response.Errors))
// 		return
// 	}
// 	j, _ := json.MarshalIndent(response, "", " ")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(j)
// }
