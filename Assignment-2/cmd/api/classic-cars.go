package main

import (
	"assignment2.ualikhan.net/internal/data"
	"assignment2.ualikhan.net/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createClassicCarsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string    `json:"name"`
		Year        int32     `json:"year"`
		Cost        data.Cost `json:"cost"`
		Description string    `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	classiccars := &data.ClassicCars{
		Name:        input.Name,
		Year:        input.Year,
		Cost:        input.Cost,
		Description: input.Description,
	}

	v := validator.New()

	if data.ValidateClassicCars(v, classiccars); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showClassicCarsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	ClassicCars := data.ClassicCars{
		ID:          id,
		CreatedAt:   time.Now(),
		Name:        "Classic Car Lexus B12",
		Year:        1934,
		Cost:        234456,
		Description: "Classic Car made in Japan. Only 200 copies.",
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"ClassicCars": ClassicCars}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
