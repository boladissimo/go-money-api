package stocks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var stockList []Entity = []Entity{{ID: 1, Code: "TSLA34", FantasyName: "Tesla"}}

//RepositoryMock is a mock of StockRepository for test porposes
type RepositoryMock struct{}

//GetAll return all stocks
func (r RepositoryMock) GetAll() []Entity {
	return stockList
}

func TestGetAll(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBody, _ := json.Marshal(stockList)

	stockController := NewController(RepositoryMock{})

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
