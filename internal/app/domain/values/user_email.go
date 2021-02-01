package values

import "errors"

// UserEmail ユーザーのメールアドレス型
type UserEmail struct {
	value string
}

// NewUserEmail メールアドレスを生成する
func NewUserEmail(val string) (*UserEmail, error) {
	if len(val) == 0 {
		return nil, errors.New("メールアドレスを指定して下さい")
	}
	return &UserEmail{value: val}, nil
}
