package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetRouter return a http.Handler instance with the routes configured
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})
	return r
}
