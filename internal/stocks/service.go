package stocks

//Service orchestrate the interaction between objects to make the request sucefully.
type Service interface {
	GetAll() []Entity
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

func (s service) GetAll() []Entity {
	return s.repository.GetAll()
}

func (s service) Create(dto DTO) Entity {
	id := s.repository.Create(dto)
	return Entity{ID: id, Code: dto.Code, FantasyName: dto.FantasyName}
}

func (s service) Delete(id int64) (int64, error) {
	return s.repository.Delete(id)
}
