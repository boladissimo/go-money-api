package interfaces_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boladissimo/go-money-api/internal/interfaces"
)

type StockControllerMock struct{}

func (s StockControllerMock) GetAll(w http.ResponseWriter, r *http.Request)  {}
func (s StockControllerMock) Create(w http.ResponseWriter, r *http.Request)  {}
func (s StockControllerMock) Delete(w http.ResponseWriter, r *http.Request)  {}
func (s StockControllerMock) GetById(w http.ResponseWriter, r *http.Request) {}

func TestHealthCheck(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBody := "ok"

	router := interfaces.GetRouter(StockControllerMock{})

	responseRecord := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)

	router.ServeHTTP(responseRecord, req)

	if expectedStatusCode != responseRecord.Code {
		t.Errorf("Expected reponse code %d. Got %d", expectedStatusCode, responseRecord.Code)
	}

	if expectedResponseBody != responseRecord.Body.String() {
		t.Errorf("Expected body with %s. Got %s", expectedResponseBody, responseRecord.Body.String())
	}
}
