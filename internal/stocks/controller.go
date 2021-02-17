package stocks

import (
	"encoding/json"
	"net/http"
)

//Controller is an interface to handle the http request to the stock domain
type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
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
