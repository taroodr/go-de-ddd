package values

import "errors"

// CircleID サークルID型
type CircleID struct {
	value string
}

// NewCircleID 新しいサークルIDを生成する
func NewCircleID(val string) (*CircleID, error) {
	if len(val) == 0 {
		return nil, errors.New("サークルIDを指定してください")
	}
	return &CircleID{value: val}, nil
}

// Value サークルIDの値を取得する
func (c *CircleID) Value() string {
	return c.value
}
