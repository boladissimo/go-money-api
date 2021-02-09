package controller

import (
	"encoding/json"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/infrastructure/repository"
)

//StockController is an interface to handle the http request to the stock domain
type StockController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

//StockControllerImpl is the main implementation of StockController
type StockControllerImpl struct {
	Repo repository.StockRepository
}

//GetAll return all stocks
func (c StockControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.Repo.GetAll())
}