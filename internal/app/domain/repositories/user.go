package repositories

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// UserRepository ユーザーリポジトリー
type UserRepository interface {
	Find(values.UserID) (*entities.User, error)
	FindByEmail(values.UserEmail) (*entities.User, error)
	FindByName(values.UserName) (*entities.User, error)
	Save(*entities.User) error
	Delete(*entities.User) error
}
