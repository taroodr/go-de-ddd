package cmd

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
	"github.com/shimar/go-de-ddd/internal/app/services"
)

// Client クライアント
type Client struct {
	userApplicationService *services.UserApplicationService
}

// ChangeName 指定したidを持つユーザーの名前を変更する
func (c *Client) ChangeName(id, name string) error {
	user, err := c.userApplicationService.Get(id)
	if err != nil {
		return err
	}

	userName, err := values.NewUserName(name)
	if err != nil {
		return err
	}
	user.ChangeName(*userName)
	return nil
}
