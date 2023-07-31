package application

import (
	"github.com/google/uuid"
	"spentio-server/internal/spentio/domain"
	"spentio-server/internal/spentio/usecases"
)

type UserService struct {
	Repository *usecases.UserRepository
}

func (u UserService) RetrieveMany(limit, offset int) (int, []*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) RetrieveOne(uuid uuid.UUID) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) CreateOne() (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) UpdateOne() (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteOne() error {
	//TODO implement me
	panic("implement me")
}
