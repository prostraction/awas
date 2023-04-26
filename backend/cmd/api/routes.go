package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/number/:id", app.getNumber)
	router.HandlerFunc(http.MethodGet, "/v1/numbers", app.getNumbersAll)
	return router
}
