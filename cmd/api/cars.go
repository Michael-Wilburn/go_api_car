package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Michael-Wilburn/go_api_car/internal/data"
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
		app.notFoundResponse(w, r)
		return
	}

	/* Create a new instance of the Car struct, containing the ID we extract from the URL
	and some dummy data. Also notice that we deliberately haven't set a  value for the Year field.*/
	car := data.Car{
		ID:         id,
		Online:     false,
		CarType:    "SUV",
		Brand:      "AUDI",
		Model:      "Q5 225HP STRONIC",
		Year:       2013,
		Kilometers: 104000,
		CarDomain:  "MBB885",
		Price:      26000000,
		InfoPrice:  24000000,
		Currency:   "$",
		ChasisCode: "JDJHAH1233DHHJJJADSJD",
		MotorCode:  "SDADDA1223",
		CreatedAt:  time.Now(),
	}

	/* Create an envelope {"car":car} instance and pass it to writeJSON(), instead
	of passing the plain car struct*/
	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"car": car}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
