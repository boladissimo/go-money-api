package stocks

import (
	"database/sql"

	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
)

//Repository is an interface to interect with the Stock database
type Repository interface {
	//GetAll return all stocks
	GetAll() []Entity
	//GetById return a stock that matches given id, if there is no match, a non null error will be returned
	GetById(id int64) (Entity, error)
	//Create inserts a stock and returns its database id
	Create(dto DTO) int64
	//Delete removes a stock that matches given id and return the number of rows affected, if there is no match, a non null error will be returned
	Delete(id int64) (int64, error)
	//Replace replaces the values from the stock given id and return the number of rows affected, if there is no match, a non null error will be returned
	Replace(entity Entity) (int64, error)
}

//repository is the main implementation of StockRepository
type repository struct {
	db *sql.DB
}

//NewRepository return a new stock repository stance //TODO: make it singleton
func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

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

func (r repository) Delete(id int64) (rowsNumber int64, err error) {
	stmt, err := r.db.Prepare("DELETE FROM stock WHERE id = ?")
	if err != nil {
		util.LogError(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		util.LogError(err)
	}

	rowsNumber, err = result.RowsAffected()

	return rowsNumber, err
}

func (r repository) GetById(id int64) (entity Entity, err error) {
	stmt := r.db.QueryRow("SELECT id, code, fantasty_name FROM stock WHERE id = ?", id)

	err = stmt.Scan(&entity.ID, &entity.Code, &entity.FantasyName)

	if err != nil {
		util.LogError(err)
	}

	return
}

func (r repository) Replace(entity Entity) (rowsNumber int64, err error) {
	stmt, err := r.db.Prepare("UPDATE stock SET code=?, fantasty_name=? WHERE id = ?")
	if err != nil {
		util.LogError(err)
	}

	result, err := stmt.Exec(entity.Code, entity.FantasyName, entity.ID)

	rowsNumber, err = result.RowsAffected()

	return
}
