package services

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/repositories"
)

// CircleService サークルサービス
type CircleService struct {
	circleRepository repositories.CircleRepository
}

// NewCircleService 新しいサークルサービスを生成する
func NewCircleService(r repositories.CircleRepository) *CircleService {
	return &CircleService{circleRepository: r}
}

// Exists 指定したサークルが存在するか否か?
func (s *CircleService) Exists(circle *entities.Circle) (bool, error) {
	c, err := s.circleRepository.FindByName(circle.Name())
	if err != nil {
		return false, err
	}
	return (c != nil), nil
}
