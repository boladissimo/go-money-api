package stocks

import (
	"database/sql"

	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
)

//Repository is an interface to interect with the Stock database
type Repository interface {
	GetAll() []Entity
	Create(dto DTO) int64
}

//repository is the main implementation of StockRepository
type repository struct {
	db *sql.DB
}

//NewRepository return a new stock repository stance //TODO: make it singleton
func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

//GetAll return all stocks TODO: remove mock
func (r repository) GetAll() []Entity {
	rows, err := r.db.Query("SELECT * FROM stock")
	if err != nil {
		panic(err)
	}

	var stocks []Entity

	for rows.Next() {
		var stock Entity

		err = rows.Scan(&stock.ID, &stock.Code, &stock.FantasyName)
		if err != nil {
			panic(err)
		}

		stocks = append(stocks, stock)
	}

	return stocks
}

func (r repository) Create(dto DTO) (id int64) {
	errr := r.db.Ping()
	if errr != nil {
		util.LogError(errr)
	}
	stmt, err := r.db.Prepare("INSERT INTO stock (code, fantasty_name) VALUES (?, ?)")
	if err != nil {
		util.LogError(err)
	}

	result, err := stmt.Exec(dto.Code, dto.FantasyName)
	if err != nil {
		panic(err)
	}

	id, _ = result.LastInsertId()
	return
}
