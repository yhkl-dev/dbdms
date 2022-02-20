package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/graphql", app.handle(nil, Register()))
	return app.enableCORS(router)
}
