package interfaces

import (
	"fmt"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/stocks"
	"github.com/gorilla/mux"
)

//GetRouter return a http.Handler instance with the routes configured
func GetRouter(stockController stocks.Controller) *mux.Router {
	r := mux.NewRouter()
	r.Use(contentTypeJSONMiddleware)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})

	r.HandleFunc("/stock", stockController.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/stock", stockController.Create).Methods(http.MethodPost)
	r.HandleFunc("/stock/{id:[0-9]+}", stockController.Delete).Methods(http.MethodDelete)
	return r
}

//contentTypeJSONMiddleware set content-type header to application/json
func contentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
