package main

import (
	"encoding/json"
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

// Define an envelope type.
type envelope map[string]interface{}

/*
Define a writeJSON() helper for sending responses. This takes the destination
http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
header map containing any additional HTTP headers we want to include in the response.
*/
// Change the data parameter to have the type envelope instaed of interface{}
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	/*Encode the data to JSON, returning the error if there was one.
	Use the json.MarshalIndent() function so that whitespace is added to the encoded
	JSON. Here we use no line prefix("") and tab indents("\t") for each element.*/
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')
	/* At this point, we know that we won't encounter any more errors before writing the
	response, so it's safe to add any headers that we want to include. We loop
	through the header map and add each header to the http.ResponseWriter header map.
	Note that it's OK if the provided header map is nil. Go doesn't throw an error
	if you try to range over (or generally, read from) a nil map.*/
	for key, value := range headers {
		w.Header()[key] = value
	}
	// Add the "Content-Type: application/json" header, then write the status code and
	// JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
