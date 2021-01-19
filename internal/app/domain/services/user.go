package services

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/repositories"
)

// UserService ユーザーサービス
type UserService struct {
	userRepository repositories.UserRepository
}

// NewUserService 新しいユーザーサービスを生成する
func NewUserService(ur repositories.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

// Exists 指定したユーザーが存在するか否か
func (s *UserService) Exists(user *entities.User) (bool, error) {
	u, err := s.userRepository.FindByName(user.Name())
	if err != nil {
		return false, err
	}
	return (u != nil), nil
}
