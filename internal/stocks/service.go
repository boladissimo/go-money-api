package stocks

//Service orchestrate the interaction between objects to make the request sucefully.
type Service interface {
	GetAll() []Entity
	GetById(id int64) (Entity, error)
	Create(dto DTO) Entity
	Delete(id int64) (int64, error)
}

type service struct {
	repository Repository
}

//NewService return a new stock service stance //TODO: make it singleton
func NewService(repository Repository) Service {
	return service{repository: repository}
}

//GetAll return all stocks
func (s service) GetAll() []Entity {
	return s.repository.GetAll()
}

//Create create an stock and return its entity
func (s service) Create(dto DTO) Entity {
	id := s.repository.Create(dto)
	return Entity{ID: id, Code: dto.Code, FantasyName: dto.FantasyName}
}

//Delete revomes given stock by id and return the number of rows affected
func (s service) Delete(id int64) (int64, error) {
	return s.repository.Delete(id)
}

func (s service) GetById(id int64) (Entity, error) {
	return s.repository.GetById(id)
}
