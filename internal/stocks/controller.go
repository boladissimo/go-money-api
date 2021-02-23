package stocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
	"github.com/gorilla/mux"
)

//Controller is an interface to handle the http request to the stock domain
type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

//NewController return a new stock controller stance //TODO: make it singleton
func NewController(service Service) Controller {
	return controller{service}
}

//controller is the main implementation of StockController
type controller struct {
	service Service
}

//GetAll return all stocks
func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.service.GetAll())
}

//GetAll return all stocks
func (c controller) Create(w http.ResponseWriter, r *http.Request) {
	var dto DTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c.service.Create(dto))
}

//Delete delete a stock given the id
func (c controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	rowsAffected, err := c.service.Delete(id)

	if err != nil {
		util.LogError(err)
	}

	if rowsAffected != 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found")
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
