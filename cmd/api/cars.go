package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Add a createCarHandler for the "POST /v1/cars" endpoint.
func (app *application) createCarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new car")
}

// Add a showCarHandler for the "GET /v1/cars/:id" endpoint.
func (app *application) showCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id == uuid.Nil {
		http.NotFound(w, r)
		return
	}

	// Otherwise, interpolated the car ID in a placeholder response.
	fmt.Fprintf(w, "show the details of car %s\n", id)
}
