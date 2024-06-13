package car

type Service interface {
	Create() error
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (c *service) Create() error {

	return nil
}
