package main

import (
	"log"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/application/controller"
	"github.com/boladissimo/go-money-api/internal/infrastructure/repository"
	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
	"github.com/boladissimo/go-money-api/internal/interfaces"
)

func main() {
	util.LogInfo("Starting go-money-api")

	stockController := controller.NewStockController(repository.StockRepositoryImpl{})
	http.Handle("/", interfaces.GetRouter(stockController))

	util.LogInfo("Serving at 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
