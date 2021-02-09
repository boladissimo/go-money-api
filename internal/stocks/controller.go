package stocks

import (
	"encoding/json"
	"net/http"
)

//Controller is an interface to handle the http request to the stock domain
type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

//NewController return a new stock controller stance //TODO: make it singleton
func NewController(repository Repository) Controller {
	return controller{repository}
}

//controller is the main implementation of StockController
type controller struct {
	repository Repository
}

//GetAll return all stocks
func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.repository.GetAll())
}
