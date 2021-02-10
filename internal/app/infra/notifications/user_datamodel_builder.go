package notifications

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/values"
	"github.com/shimar/go-de-ddd/internal/app/infra/models"
)

type UserDataModelBuilder struct {
	userID values.UserID
	userName values.UserName
}

func (u *UserDataModelBuilder) Name(userName values.UserName) {
	u.userName = userName
}

func (u *UserDataModelBuilder) ID(userID values.UserID) {
	u.userID = userID
}

func (u *UserDataModelBuilder) Build() *models.UserDataModel {
	return &models.UserDataModel{
		ID:   u.userID.Value(),
		Name: u.userName.Value(),
	}
}