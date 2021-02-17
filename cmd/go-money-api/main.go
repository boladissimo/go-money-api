package main

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	util.LogInfo("Serving at " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
