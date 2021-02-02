package main

import (
	"log"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/interfaces"
	"github.com/boladissimo/go-money-api/internal/util"
)

func main() {
	util.LogInfo("Starting go-money-api")
	http.Handle("/", interfaces.GetRouter())
	util.LogInfo("Serving at 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
