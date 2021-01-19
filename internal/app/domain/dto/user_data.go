package dto

import "github.com/shimar/go-de-ddd/internal/app/domain/entities"

// UserData ユーザー情報DTO
type UserData struct {
	id   string
	name string
}

// NewUserData 新しいユーザー情報DTOを生成する
func NewUserData(user *entities.User) *UserData {
	id := user.ID()
	name := user.Name()
	return &UserData{
		id:   id.Value(),
		name: name.Value(),
	}
}

// ID ユーザーIDを取得する
func (u *UserData) ID() string {
	return u.id
}

// Name ユーザー名を取得する
func (u *UserData) Name() string {
	return u.name
}
