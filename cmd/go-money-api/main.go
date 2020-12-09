package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/util"

	"github.com/gorilla/mux"
)

func main() {
	util.LogInfo("Starting go-money-api")
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})

	http.Handle("/", r)
	util.LogInfo("Serving at 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
