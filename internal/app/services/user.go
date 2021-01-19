package services

import (
	"errors"

	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/repositories"
	"github.com/shimar/go-de-ddd/internal/app/domain/services"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// UserApplicationService ユーザーアプリケーションサービス
type UserApplicationService struct {
	userRepository repositories.UserRepository
	userService    *services.UserService
}

// NewUserApplicationService 新しいユーザーアプリケーションサービスを生成する
func NewUserApplicationService(r repositories.UserRepository, s *services.UserService) *UserApplicationService {
	return &UserApplicationService{userRepository: r, userService: s}
}

// Register ユーザーを登録する
func (s *UserApplicationService) Register(name string) error {
	// ユーザー名(値オブジェクト)を生成する
	userName, err := values.NewUserName(name)
	if err != nil {
		return err
	}
	// ユーザー(エンティティ)を生成する
	user, err := entities.NewUser(nil, *userName)
	if err != nil {
		return err
	}
	// 存在を検証する
	exists, err := s.userService.Exists(user)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ユーザーは既に存在しています")
	}
	// ユーザーを登録する
	return s.userRepository.Save(user)
}

// Get 指定したIDを持つユーザーを取得する
func (s *UserApplicationService) Get(id string) (*entities.User, error) {
	userID, err := values.NewUserID(id)
	if err != nil {
		return nil, err
	}
	return s.userRepository.Find(*userID)
}
