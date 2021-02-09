package stocks

//Repository is an interface to interect with the Stock database
type Repository interface {
	GetAll() []Entity
}

//repository is the main implementation of StockRepository
type repository struct{}

//NewRepository return a new stock repository stance //TODO: make it singleton
func NewRepository() Repository {
	return repository{}
}

//GetAll return all stocks TODO: remove mock
func (r repository) GetAll() []Entity {
	return []Entity{{ID: 1, Code: "TSLA34", FantasyName: "Tesla"}}
}
