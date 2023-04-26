package main

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getNumber(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
	}
	number := models.Number{
		ID:     id,
		Name:   "000000000001",
		Status: "Created",
	}
	err = app.writeJson(w, http.StatusOK, number, "number")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getNumbersAll(w http.ResponseWriter, r *http.Request) {

}
