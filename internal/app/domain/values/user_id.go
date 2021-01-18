package values

import "errors"

// UserID ユーザーID型
type UserID struct {
	value string
}

// NewUserID 新規ユーザーIDを生成する
func NewUserID(val string) (*UserID, error) {
	if len(val) == 0 {
		return nil, errors.New("無効な値だよ")
	}
	return &UserID{value: val}, nil
}

// Value 値を取得する
func (u *UserID) Value() string {
	return u.value
}
