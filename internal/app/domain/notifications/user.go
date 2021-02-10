package notifications

import "github.com/shimar/go-de-ddd/internal/app/domain/values"

type UserNotification interface {
	ID(id values.UserID)
	Name(id values.UserName)
}

