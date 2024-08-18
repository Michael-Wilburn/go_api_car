package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	/*Convert the notFoundResponse() helper to a http.Handler using the
	http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	Not Found Responses.*/
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	/*Likewise, convert tje methodNotAllowedResponse() helper to a http.Handler and set
	it as the custom error handler for 405 Method Not Allowed responses.*/
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	/*Register the relevant methods, URL patterns and handler functions for our
	endpoints using HandleFunc() method. Note that http.MethodGet and http.MethodPost
	are constants which equate to the string "GET" and "POST" respectively.*/
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/cars", app.createCarHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cars/:id", app.showCarHandler)

	// Return the httprouter instance
	return router
}
