package main

import (
	"assignment3.ualikhan.net/internal/data"
	"assignment3.ualikhan.net/internal/validator"
	"errors"
	"fmt"
	"net/http"
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

	err = app.models.ClassicCars.Insert(classiccars)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/classic-cars/%d", classiccars.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"classiccars": classiccars}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showClassicCarsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	classiccars, err := app.models.ClassicCars.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"classiccars": classiccars}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateClassicCarsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	classiccars, err := app.models.ClassicCars.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name        *string    `json:"name"`
		Year        *int32     `json:"year"`
		Cost        *data.Cost `json:"cost"`
		Description *string    `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		classiccars.Name = *input.Name
	}
	if input.Year != nil {
		classiccars.Year = *input.Year
	}
	if input.Cost != nil {
		classiccars.Cost = *input.Cost
	}
	if input.Description != nil {
		classiccars.Description = *input.Description
	}

	v := validator.New()
	if data.ValidateClassicCars(v, classiccars); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.ClassicCars.Update(classiccars)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	
	err = app.writeJSON(w, http.StatusOK, envelope{"classiccars": classiccars}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteClassicCarsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.ClassicCars.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "classic_car successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
