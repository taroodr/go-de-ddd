package repositories

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// CircleRepository サークルリポジトリー
type CircleRepository interface {
	Save(*entities.Circle) error
	Find(*values.CircleID) (*entities.Circle, error)
	FindByName(*values.CircleName) (*entities.Circle, error)
}
