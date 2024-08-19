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
	/*
		Declare an anonymous struct to hold the information that we expect to be in the
		HTTP request body (note that the field names and types in the struct are a subset
		of the Car struct that we created earlier). This struct will be our *target decode destination.
	*/
	var input struct {
		Online     bool    `json:"online"`
		CarType    string  `json:"car_type"`
		Brand      string  `json:"brand"`
		Model      string  `json:"model"`
		Year       int32   `json:"year"`
		Kilometers int64   `json:"kilometers"`
		CarDomain  string  `json:"car_domain"`
		Price      float64 `json:"price"`
		InfoPrice  float64 `json:"info_price"`
		Currency   string  `json:"currency"`
		ChasisCode string  `json:"chasis_code"`
		MotorCode  string
	}

	/*
		Use the new readJSON() helper to decode the request body into the input struct.
		If this returns an error we send the client the error message along with a 400
		Bad Request status code, just like before.
	*/
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
