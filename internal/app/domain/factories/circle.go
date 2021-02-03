package factories

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// CircleFactory サークルファクトリー
type CircleFactory interface {
	Create(*values.CircleName, *entities.User) (*entities.Circle, error)
}
