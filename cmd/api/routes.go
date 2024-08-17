package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	/*Register the relevant methods, URL patterns and handler functions for our
	endpoints using HandleFunc() method. Note that http.MethodGet and http.MethodPost
	are constants which equate to the string "GET" and "POST" respectively.*/
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/cars", app.createCarHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cars/:id", app.showCarHandler)

	// Return the httprouter instance
	return router
}
