package services

import (
	"errors"

	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/domain/repositories"
	"github.com/shimar/go-de-ddd/internal/app/domain/services"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// UserRegisterService ユーザー登録サービス
type UserRegisterService interface {
	Handle(*UserRegisterCommand) error
}

// UserRegisterServiceImpl ユーザー登録サービス実装
type UserRegisterServiceImpl struct {
	userRepository repositories.UserRepository
	userService    *services.UserService
}

// NewUserRegisterService ユーザー登録サービスを生成する
func NewUserRegisterService(r repositories.UserRepository, s *services.UserService) UserRegisterService {
	return &UserRegisterServiceImpl{userRepository: r, userService: s}
}

// Handle ユーザーを登録する
func (s *UserRegisterServiceImpl) Handle(cmd *UserRegisterCommand) error {
	// ユーザー名を生成する
	userName, err := values.NewUserName(cmd.name)
	if err != nil {
		return err
	}

	// ユーザーを生成する
	user, err := entities.NewUser(nil, *userName)
	if err != nil {
		return err
	}

	// 重複を検証する
	exists, err := s.userService.Exists(user)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ユーザーは既に存在しています")
	}

	return s.userRepository.Save(user)
}

// UserRegisterCommand ユーザー登録コマンド構造体
type UserRegisterCommand struct {
	name string
}

// NewUserRegisterCommand ユーザー登録コマンドを生成する
func NewUserRegisterCommand() *UserRegisterCommand {
	return &UserRegisterCommand{}
}

// SetName 名前を設定する
func (c *UserRegisterCommand) SetName(s string) {
	c.name = s
}
