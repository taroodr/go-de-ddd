package entities

import (
	"errors"

	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// Circle サークル型
type Circle struct {
	id      *values.CircleID
	name    *values.CircleName
	owner   *User
	members []*User
}

// NewCircle サークルを生成する
func NewCircle(id *values.CircleID, name *values.CircleName, owner *User, members []*User) (*Circle, error) {
	if id == nil {
		return nil, errors.New("id is nil")
	}
	if name == nil {
		return nil, errors.New("name is nil")
	}
	if owner == nil {
		return nil, errors.New("owner is nil")
	}
	if members == nil {
		return nil, errors.New("members is nil")
	}
	return &Circle{id: id, name: name, owner: owner, members: members}, nil
}

// ID サークルIDを取得する
func (c *Circle) ID() *values.CircleID {
	return c.id
}

// Name サークル名を取得する
func (c *Circle) Name() *values.CircleName {
	return c.name
}

// Owner オーナーを取得する
func (c *Circle) Owner() *User {
	return c.owner
}

// Member メンバーを取得する
func (c *Circle) Members() []*User {
	return c.members
}

func (c *Circle) IsFull() bool {
	return len(c.members) >= 29
}

// AddMember メンバーを追加する
func (c *Circle) Join(u *User) error {
	if c.IsFull() {
		return errors.New("このサークルは満員です")
	}
	c.members = append(c.members, u)
	return nil
}
