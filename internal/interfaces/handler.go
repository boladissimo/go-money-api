package interfaces

import (
	"fmt"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/application/controller"
	"github.com/gorilla/mux"
)

//GetRouter return a http.Handler instance with the routes configured
func GetRouter(stockController controller.StockController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})

	r.HandleFunc("/stock", stockController.GetAll).Methods(http.MethodGet)
	return r
}