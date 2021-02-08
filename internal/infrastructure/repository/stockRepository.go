package repository

import "github.com/boladissimo/go-money-api/internal/domain/stock"

//StockRepository is an interface to interect with the Stock database
type StockRepository interface {
	GetAll() []stock.Entity
}

//StockRepositoryImpl is the main implementation of StockRepository
type StockRepositoryImpl struct{}

//GetAll return all stocks TODO: remove mock
func (r StockRepositoryImpl) GetAll() []stock.Entity {
	return []stock.Entity{{ID: 1, Code: "TSLA34", FantasyName: "Tesla"}}
}
