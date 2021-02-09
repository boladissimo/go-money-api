package interfaces

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boladissimo/go-money-api/internal/application/controller"
	"github.com/boladissimo/go-money-api/internal/infrastructure/repository"
)

func TestHealthCheck(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedResponseBody := "ok"

	stockController := controller.StockControllerImpl{Repo: repository.StockRepositoryImpl{}}
	router := GetRouter(stockController)

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
