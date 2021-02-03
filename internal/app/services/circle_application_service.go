package services

import (
	"errors"

	"github.com/shimar/go-de-ddd/internal/app/domain/factories"
	"github.com/shimar/go-de-ddd/internal/app/domain/repositories"
	"github.com/shimar/go-de-ddd/internal/app/domain/services"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// CircleApplicationService サークルアプリケーションサービス
type CircleApplicationService struct {
	circleFactory    factories.CircleFactory
	circleRepository repositories.CircleRepository
	circleService    services.CircleService
	userRepository   repositories.UserRepository
}

// NewCircleApplicationService サークルアプリケーションサービスを生成する
func NewCircleApplicationService(cf factories.CircleFactory, cr repositories.CircleRepository, cs services.CircleService, ur repositories.UserRepository) *CircleApplicationService {
	return &CircleApplicationService{circleFactory: cf, circleRepository: cr, circleService: cs, userRepository: ur}
}

// Create サークルを登録する
func (s *CircleApplicationService) Create(cmd *CircleCreateCommand) error {
	// NOTE: Tx Begin
	// オーナーとなるユーザーを取得
	oID, err := values.NewUserID(cmd.UserID())
	if err != nil {
		return err
	}
	owner, err := s.userRepository.Find(*oID)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("サークルのオーナーとなるユーザーが見つかりません")
	}

	// サークルの存在チェック
	name, err := values.NewCircleName(cmd.Name())
	if err != nil {
		return err
	}
	circle, err := s.circleFactory.Create(name, owner)
	if err != nil {
		return err
	}
	exists, err := s.circleService.Exists(circle)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("サークルはすでに存在しています")
	}

	// サークルを登録
	if err := s.circleRepository.Save(circle); err != nil {
		return err
	}
	// NOTE: Tx Commit
	return nil
}

// CircleCreateCommand サークル生成コマンド
type CircleCreateCommand struct {
	userID string
	name   string
}

// UserID ユーザーIDを取得する
func (c *CircleCreateCommand) UserID() string {
	return c.userID
}

// Name サークル名を取得する
func (c *CircleCreateCommand) Name() string {
	return c.name
}
