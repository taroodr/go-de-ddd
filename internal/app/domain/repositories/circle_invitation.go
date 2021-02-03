package repositories

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
)

// CircleInvitationRepository サークル勧誘リポジトリー
type CircleInvitationRepository interface {
	Save(*entities.CircleInvitation) error
}
