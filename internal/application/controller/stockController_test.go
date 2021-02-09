package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boladissimo/go-money-api/internal/domain/stock"
)

var stockList []stock.Entity = []stock.Entity{{ID: 1, Code: "TSLA34", FantasyName: "Tesla"}}

//StockRepositoryMock is a mock of StockRepository for test porposes
type StockRepositoryMock struct{}

//GetAll return all stocks
func (r StockRepositoryMock) GetAll() []stock.Entity {
	return stockList
}

func TestGetAll(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBody, _ := json.Marshal(stockList)

	stockController := NewStockController(StockRepositoryMock{})

	responseRecord := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/stock", nil)

	stockController.GetAll(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if bytes.Equal(expectedResponseBody, responseRecord.Body.Bytes()) {
		t.Errorf("Expected body with %s. Got %s", string(expectedResponseBody), responseRecord.Body.String())
	}
}
