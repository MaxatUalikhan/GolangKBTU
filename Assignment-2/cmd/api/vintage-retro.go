package main

import (
	"assignment2.ualikhan.net/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createVintageRetroHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Year int32  `json:"year"`
		Cost int32  `json:"cost"`
		Type string `json:"type"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showVintageRetroHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	VintageRetroItem := data.VintageRetro{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Classic Car Lexus B12",
		Year:      1934,
		Cost:      234456,
		Type:      "Classic Car",
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"VintageRetroItem": VintageRetroItem}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
