package services

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/repositories"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// UserDeleteService ユーザー削除サービス
type UserDeleteService struct {
	userRepository repositories.UserRepository
}

// NewUserDeleteService ユーザー削除サービスを生成する
func NewUserDeleteService(r repositories.UserRepository) *UserDeleteService {
	return &UserDeleteService{userRepository: r}
}

// Handle ユーザーを削除する
func (s *UserDeleteService) Handle(cmd *UserDeleteCommand) error {
	userID, err := values.NewUserID(cmd.id)
	if err != nil {
		return err
	}

	user, err := s.userRepository.Find(*userID)
	if err != nil {
		return err
	}

	if user == nil {
		return nil
	}

	return s.userRepository.Delete(user)
}

// UserDeleteCommand ユーザー削除コマンド構造体
type UserDeleteCommand struct {
	id string
}

// NewUserDeleteCommand ユーザー削除コマンドを生成する
func NewUserDeleteCommand(id string) *UserDeleteCommand {
	return &UserDeleteCommand{id: id}
}
