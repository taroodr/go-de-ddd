package entities

import (
	"github.com/google/uuid"
	"github.com/shimar/go-de-ddd/internal/app/domain/notifications"
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// User ユーザー型
type User struct {
	id   values.UserID
	name values.UserName
}

// NewUser ユーザーを生成する
func NewUser(id *values.UserID, name values.UserName) (*User, error) {
	if id == nil {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		id, err = values.NewUserID(uuid.String())
		if err != nil {
			return nil, err
		}
	}
	return &User{id: *id, name: name}, nil
}

// ID ユーザーのIDを取得する
func (u *User) ID() values.UserID {
	return u.id
}

// Name ユーザーの名前を取得する
func (u *User) Name() values.UserName {
	return u.name
}

// ChangeName ユーザーの名前を変更する
func (u *User) ChangeName(name values.UserName) {
	u.name = name
}

func (u *User) Notify(note notifications.UserNotification) {
	note.ID(u.id)
	note.Name(u.name)
}