package entities

// CircleInvitation サークル勧誘型
type CircleInvitation struct {
	circle *Circle
	from   *User
	to     *User
}

// NewCircleInvitation サークル勧誘を生成する
func NewCircleInvitation(circle *Circle, from *User, to *User) *CircleInvitation {
	return &CircleInvitation{circle: circle, from: from, to: to}
}
