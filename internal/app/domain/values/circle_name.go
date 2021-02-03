package values

import "errors"

// CircleName サークル名型
type CircleName struct {
	value string
}

// NewCircleName 新しいサークル名を生成する
func NewCircleName(val string) (*CircleName, error) {
	if len(val) == 0 {
		return nil, errors.New("サークル名を指定してください")
	}
	if len(val) < 3 {
		return nil, errors.New("サークル名は3文字以上です")
	}
	if len(val) > 20 {
		return nil, errors.New("サークル名は20文字以下です")
	}
	return &CircleName{value: val}, nil
}

// Value サークル名の値を取得する
func (c *CircleName) Value() string {
	return c.value
}

// Equals 同じサークル名か否か?
func (c *CircleName) Equals(other *CircleName) bool {
	return c.value == other.value
}
