package stocks_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
	"github.com/boladissimo/go-money-api/internal/interfaces"
	"github.com/boladissimo/go-money-api/internal/stocks"
)

var stockList []stocks.Entity = []stocks.Entity{{ID: 1, Code: "TSLA34", FantasyName: "Tesla"}}
var stockEntity stocks.Entity = stocks.Entity{ID: 1, Code: "foo", FantasyName: "bar"}

//ServiceMock is a mock of StockService for test porposes
type ServiceMock struct{}

//GetAll return all stocks
func (r ServiceMock) GetAll() []stocks.Entity {
	return stockList
}

func (r ServiceMock) Create(dto stocks.DTO) (entity stocks.Entity) {
	entity = stockEntity
	return
}

func (r ServiceMock) Delete(id int64) (int64, error) {
	util.LogInfo("given id is ", id)
	return id, nil
}

func TestGetAll_stocksPresent_200AndStockList(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBodyBytes, _ := json.Marshal(stockList)
	expectedResponseBody := string(expectedResponseBodyBytes) + "\n"

	stockController := stocks.NewController(ServiceMock{})

	responseRecord := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/stock", nil)

	stockController.GetAll(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if expectedResponseBody != responseRecord.Body.String() {
		t.Errorf("Expected body with %s. Got %s", expectedResponseBody, responseRecord.Body.String())
	}
}

func TestCreate_validPayload_201AndCreatedStockEntity(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedResponseBodyBytes, _ := json.Marshal(stockEntity)
	expectedResponseBody := string(expectedResponseBodyBytes) + "\n"

	stockController := stocks.NewController(ServiceMock{})

	responseRecord := httptest.NewRecorder()

	requestBody, _ := json.Marshal(stocks.DTO{Code: stockEntity.Code, FantasyName: stockEntity.FantasyName})
	req, _ := http.NewRequest(http.MethodPost, "/stock", bytes.NewReader(requestBody))

	stockController.Create(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if expectedResponseBody != responseRecord.Body.String() {
		t.Errorf("Expected body with %s. Got %s", expectedResponseBody, responseRecord.Body.String())
	}
}

func TestDelete_validStockId_204AndNoContent(t *testing.T) {
	expectedStatusCode := http.StatusNoContent

	stockController := stocks.NewController(ServiceMock{})
	router := interfaces.GetRouter(stockController)

	responseRecord := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/stock/1", nil)

	router.ServeHTTP(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if nil != responseRecord.Body.Bytes() {
		t.Errorf("Expected body to be nil. Got %s", responseRecord.Body.String())
	}
}

func TestDelete_invalidStockId_404AndNotFoundMessage(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedResponseBody := []byte("not found")

	stockController := stocks.NewController(ServiceMock{})
	router := interfaces.GetRouter(stockController)

	responseRecord := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/stock/2", nil)

	router.ServeHTTP(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if string(expectedResponseBody) != responseRecord.Body.String() {
		t.Errorf("Expected body with %s. Got %s", string(expectedResponseBody), responseRecord.Body.String())
	}
}
