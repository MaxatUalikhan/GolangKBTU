package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/classic-cars", app.createClassicCarsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/classic-cars/:id", app.showClassicCarsHandler)
	router.HandlerFunc(http.MethodPut, "/v1/classic-cars/:id", app.updateClassicCarsHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/classic-cars/:id", app.deleteClassicCarsHandler)

	return router
}
