package main

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

/*
Retrive the "id" URL parameter from the current request context, then convert it to a uuid
and return it. If the operation isn't successful, return 0 and err.
*/
func (app *application) readIDParam(r *http.Request) (uuid.UUID, error) {
	/* When httprouter is parsing a request, any interpolated URL parameters will be
	stored in the request context. We can use the ParamsFromContext() function to
	retrieve a slice containing these parameter name and values.*/
	params := httprouter.ParamsFromContext(r.Context())

	id, err := uuid.Parse(params.ByName("id"))
	if err != nil {
		return uuid.Nil, errors.New("invalid id parameter")
	}

	return id, nil
}
