package main

import (
	"log"
	"net/http"

	"github.com/boladissimo/go-money-api/internal/infrastructure/config/db"
	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
	"github.com/boladissimo/go-money-api/internal/interfaces"
	"github.com/boladissimo/go-money-api/internal/stocks"
)

func main() {
	util.LogInfo("Starting go-money-api")

	stockRepository := stocks.NewRepository(db.GetDB())
	stockService := stocks.NewService(stockRepository)
	stockController := stocks.NewController(stockService)
	http.Handle("/", interfaces.GetRouter(stockController))

	util.LogInfo("Serving at 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
