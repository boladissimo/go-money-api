package stocks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var stockList []Entity = []Entity{{ID: 1, Code: "TSLA34", FantasyName: "Tesla"}}
var stockEntity Entity = Entity{ID: 1, Code: "foo", FantasyName: "bar"}

//ServiceMock is a mock of StockService for test porposes
type ServiceMock struct{}

//GetAll return all stocks
func (r ServiceMock) GetAll() []Entity {
	return stockList
}

func (r ServiceMock) Create(dto DTO) (entity Entity) {
	entity = stockEntity
	return
}
func TestGetAll(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBody, _ := json.Marshal(stockList)

	stockController := NewController(ServiceMock{})

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

func TestCreate(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedResponseBody, _ := json.Marshal(stockEntity)

	stockController := NewController(ServiceMock{})

	responseRecord := httptest.NewRecorder()

	requestBody, _ := json.Marshal(DTO{Code: stockEntity.Code, FantasyName: stockEntity.FantasyName})
	req, _ := http.NewRequest(http.MethodPost, "/stock", bytes.NewReader(requestBody))

	stockController.Create(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if bytes.Equal(expectedResponseBody, responseRecord.Body.Bytes()) {
		t.Errorf("Expected body with %s. Got %s", string(expectedResponseBody), responseRecord.Body.String())
	}
}
