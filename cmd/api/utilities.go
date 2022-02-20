package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/imdario/mergo"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "applicatin/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}
	app.writeJSON(w, statusCode, theError, "error")
}

func merge(dest graphql.Fields, src ...graphql.Fields) (result graphql.Fields) {
	for _, data := range src {
		if err := mergo.Merge(&dest, data); err != nil {
			panic(err)
		}
	}
	return dest
}
