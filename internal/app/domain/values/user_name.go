package values

import "errors"

// UserName ユーザー名型
type UserName struct {
	value string
}

// NewUserName ユーザー名を生成する
func NewUserName(val string) (*UserName, error) {
	if len(val) < 3 {
		return nil, errors.New("ユーザー名は3文字以上です。")
	}
	if len(val) > 20 {
		return nil, errors.New("ユーザー名は20文字以上です。")
	}
	return &UserName{value: val}, nil
}

// Value 値を取得する
func (u *UserName) Value() string {
	return u.value
}
