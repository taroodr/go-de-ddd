package cmd

import (
	"github.com/shimar/go-de-ddd/internal/app/services"
)

// Client クライアント
type Client struct {
	userRegisterService services.UserRegisterService
}

// Register ユーザーを登録する
func (c *Client) Register(name string) error {
	cmd := services.NewUserRegisterCommand()
	cmd.SetName(name)

	return c.userRegisterService.Handle(cmd)
}
