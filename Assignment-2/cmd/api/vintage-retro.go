package main

import (
	"assignment2.ualikhan.net/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createVintageRetroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new vintage-retro item")
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
