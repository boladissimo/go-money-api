package interfaces

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StockControllerMock struct{}

func (s StockControllerMock) GetAll(w http.ResponseWriter, r *http.Request) {}
func (s StockControllerMock) Create(w http.ResponseWriter, r *http.Request) {}
func (s StockControllerMock) Delete(w http.ResponseWriter, r *http.Request) {}

func TestHealthCheck(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBody := "ok"

	router := GetRouter(StockControllerMock{})

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
