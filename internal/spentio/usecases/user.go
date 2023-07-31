package usecases

import (
	"github.com/google/uuid"
	"spentio-server/internal/spentio/domain"
)

type UserRepository interface {
	RetrieveMany(limit, offset int) (int, []*domain.User, error)
	RetrieveOne(uuid uuid.UUID) (*domain.User, error)
	CreateOne() (*domain.User, error)
	UpdateOne() (*domain.User, error)
	DeleteOne() error
}
