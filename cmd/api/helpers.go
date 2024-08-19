package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Decode the request body into the target destination.
	// Initialize a new json.Decoder instance which reads from the request body, and
	// then use the Decode() method to decode the body contents into the input struct.
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		// If there is an error during decoding, start the triage ...
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		/* Use the errors.As() function to check whether the error has the type *json.SyntaxtError.
		If it does, then return a plain-english error message
		which includes the location of the problem. */
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formated JSON (at character %d)", syntaxError.Offset)

		/* In some circumstances Decode() may also return an io.ErrUnexpectedEOF error
		for syntax errors in the JSON. So we check for this using errors.Is() and return a
		generic error message.*/
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formated JSON")

		/* Likewise, catch any *json.UnmarshalTypeError errors. These occur when the
		 		JSON value is the wrong type for the target destination. If the error relates
				to a specific field, then we include that in our error message to make it
		 		easier for the client to debug.*/
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		/* An io.EOF error will be returned by Decode() if the request body is empty. We check
		for this with errors.Is() and return a plain-english error message instead.*/
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		/*
			A json.InvalidUnmarshalError error will be returned if we pass a non-nil
			pointer to Decode(). We catch this and panic, rather than returning an error
			to our handler. At the end of this chapter we'll talk about panicking
			versus returning errors, and discuss why it's an appropriate thing to do in
			this specific situation.
		*/
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		// For anything else, return the error message as-is.
		default:
			return err
		}
	}
	return nil
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
